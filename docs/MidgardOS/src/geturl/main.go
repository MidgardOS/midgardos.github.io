package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jlaffaye/ftp"
)

// ProgressWriter wraps an io.Writer to track and display download progress.
type ProgressWriter struct {
	writer    io.Writer
	total     int64
	written   int64
	lastPrint time.Time
	startTime time.Time
	mu        sync.Mutex
}

// NewProgressWriter creates a new ProgressWriter with the given total size.
// If total is 0 or negative, percentage will not be displayed.
func NewProgressWriter(writer io.Writer, total int64) *ProgressWriter {
	return &ProgressWriter{
		writer:    writer,
		total:     total,
		startTime: time.Now(),
		lastPrint: time.Now(),
	}
}

// Write implements io.Writer and tracks progress.
func (pw *ProgressWriter) Write(p []byte) (int, error) {
	n, err := pw.writer.Write(p)
	if n > 0 {
		pw.mu.Lock()
		pw.written += int64(n)
		// Update display every 100ms to avoid spamming the terminal
		if time.Since(pw.lastPrint) >= 100*time.Millisecond {
			pw.printProgress()
			pw.lastPrint = time.Now()
		}
		pw.mu.Unlock()
	}
	return n, err
}

// printProgress displays the current download progress.
func (pw *ProgressWriter) printProgress() {
	elapsed := time.Since(pw.startTime)
	speed := float64(pw.written) / elapsed.Seconds()

	if pw.total > 0 {
		percent := float64(pw.written) / float64(pw.total) * 100
		// ANSI terminal codes: \033[2K clears the current line, and \r moves the cursor back to the start of the line
		fmt.Fprintf(os.Stderr, "\033[2K\rDownloading: %.2f%% (%s / %s) - %s/s",
			percent,
			formatBytes(pw.written),
			formatBytes(pw.total),
			formatBytes(int64(speed)))
	} else {
		fmt.Fprintf(os.Stderr, "\033[2K\rDownloading: %s - %s/s",
			formatBytes(pw.written),
			formatBytes(int64(speed)))
	}
}

// Finish prints the final progress line with a newline.
func (pw *ProgressWriter) Finish() {
	pw.mu.Lock()
	defer pw.mu.Unlock()
	pw.printProgress()
	fmt.Fprintln(os.Stderr)
}

// formatBytes formats a byte count into a human-readable string.
func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %ciB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// showVersion prints the version information of the geturl tool.
func showVersion() {
	println("Version: 1.0.0")
	println("Copyright (c) 2026 YggdrasilSoft, LLC.")
	println("License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>")
	println("This is free software: you are free to change and redistribute it. There is NO ")
	println("WARRANTY, to the extent permitted by law.\n")
	println("Author: Gary L. Greene Jr.")
}

// showHelp prints the usage information and command-line options for the geturl tool.
// It displays supported protocols (HTTP, HTTPS, FTP) and example usage.
func showHelp() {
	println("geturl - a simple file download tool")
	println("Usage: geturl [options] <url>\n")
	println("Options:")
	println("  -v|--version        Show version information")
	println("  -h|--help           Show this help message")
	println("  -u|--url <url>      Specify the URL to download (supports HTTP, HTTPS, FTP)")
	println("  -o|--output <file>  Specify the output file name")
	println("  -s|--secure         Do TLS/SSL validation")
	println("  -O|--remote-name    Use the remote file name. Must be in the last position")
	println("                      after the URL")
	println("\nExamples:")
	println("  geturl -u https://example.com/file.txt -o file.txt")
	println("  geturl -u ftp://ftp.example.com/path/file.tar.gz -o file.tar.gz")
	println("  geturl -u https://example.com/file.txt -O\n")
	showVersion()
}

// processArgs parses command-line arguments and returns the URL, output file name, and TLS validation flag.
func processArgs(args []string) (string, string, bool) {
	var urlStr string
	var fileName string
	var useTLSValidation bool = false

	// needed to ensure we have the required arguments. Otherwise, print our help message and exit
	hasURL := false
	hasOutput := false

	// read in our arguments ARGV
	for position, arg := range args {
		if arg == "--help" || arg == "-h" {
			showHelp()
			os.Exit(0)
		}
		if arg == "--version" || arg == "-v" {
			showVersion()
			os.Exit(0)
		}
		if arg == "--url" || arg == "-u" {
			// put the next argument in the list into a variable called url
			// we need to check if there is a next argument first
			// if there isn't, we should print an error message and exit
			if position+1 < len(os.Args) {
				urlStr = os.Args[position+1]
			} else {
				println("Error: No URL specified")
				os.Exit(1)
			}
			// parse url via net/url to make sure it's valid
			if _, err := url.ParseRequestURI(urlStr); err != nil {
				println("Error: Invalid URL")
				os.Exit(1)
			}
			hasURL = true
		}
		if arg == "--output" || arg == "-o" {
			// put the next argument in the list into a variable called output
			// we need to check if there is a next argument first
			// if there isn't, we should print an error message and exit
			if position+1 < len(os.Args) {
				fileName = os.Args[position+1]
			} else {
				println("Error: No output file specified")
				os.Exit(1)
			}
			hasOutput = true
		}
		if arg == "--remote-name" || arg == "-O" {
			// if we have a url, we can use the path component of the url as the file name
			if hasURL {
				parsedURL, err := url.Parse(urlStr)
				if err != nil {
					println("Error: Invalid URL")
					os.Exit(1)
				}
				// get the last part of the path as the file name
				pathParts := strings.Split(parsedURL.Path, "/")
				fileName = pathParts[len(pathParts)-1]
				hasOutput = true
			} else {
				println("Error: --remote-name option requires a URL to be specified first")
				os.Exit(1)
			}
		}
		if arg == "--secure" || arg == "-s" {
			useTLSValidation = true
		}
	}

	// if we don't have a url, print an error message and exit
	if !hasURL {
		println("Error: No URL specified")
		showHelp()
		os.Exit(1)
	}

	// if we don't have an output file, print an error message and exit
	if !hasOutput {
		println("Error: No output file specified")
		showHelp()
		os.Exit(1)
	}

	return urlStr, fileName, useTLSValidation
}

// main is the entry point for the geturl application.
// It parses command-line arguments, validates the provided URL and output file,
// and initiates the download process.
func main() {
	var urlStr string
	var fileName string
	var useTLSValidation bool

	urlStr, fileName, useTLSValidation = processArgs(os.Args)

	println("Downloading file from URL: ", urlStr)
	println("Saving file to:            ", fileName)

	// Download the file with SSL validation disabled
	if err := downloadFile(urlStr, fileName, useTLSValidation); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to download file: %v\n", err)
		os.Exit(1)
	}

	println("Download completed successfully!")
}

// downloadFile downloads a file from the given URL to the specified path.
// It supports HTTP, HTTPS, and FTP protocols. SSL certificate verification
// is disabled for HTTPS to work on systems without root CA bundles.
// The function automatically detects the protocol from the URL scheme and
// routes to the appropriate handler.
func downloadFile(urlStr, filepath string, useTLSValidation bool) error {
	// Parse URL to determine protocol
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}

	// Route to appropriate download handler based on scheme
	scheme := strings.ToLower(parsedURL.Scheme)
	switch scheme {
	case "http", "https":
		return downloadHTTP(urlStr, filepath, useTLSValidation)
	case "ftp":
		return downloadFTP(parsedURL, filepath)
	default:
		return fmt.Errorf("unsupported protocol: %s (supported: http, https, ftp)", scheme)
	}
}

// downloadHTTP downloads a file via HTTP or HTTPS protocol.
// TLS certificate verification is disabled to support systems without
// root CA certificate bundles. The function sets appropriate timeouts
// and handles the complete request/response cycle including file creation.
func downloadHTTP(urlStr, filepath string, useTLSValidation bool) error {
	// Create HTTP client with TLS verification disabled
	// first, create a transport
	tr := &http.Transport{}
	if !useTLSValidation {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			// Set reasonable timeouts
			IdleConnTimeout:       30 * time.Second,
			ResponseHeaderTimeout: 30 * time.Second,
		}
	} else {
		tr = &http.Transport{
			// Set reasonable timeouts
			IdleConnTimeout:       30 * time.Second,
			ResponseHeaderTimeout: 30 * time.Second,
		}
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   5 * time.Minute, // Overall timeout for the request
	}

	// Create the HTTP request
	resp, err := client.Get(urlStr)
	if err != nil {
		return fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Create the output file
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	// Create progress writer with total size from Content-Length header
	progress := NewProgressWriter(out, resp.ContentLength)

	// Write the response body to file with progress tracking
	_, err = io.Copy(progress, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	// Finish progress display
	progress.Finish()

	return nil
}

// downloadFTP downloads a file via FTP protocol.
// It handles both anonymous and authenticated FTP connections.
// If no credentials are provided in the URL, anonymous login is used.
// The default FTP port (21) is used if no port is specified.
func downloadFTP(parsedURL *url.URL, filepath string) error {
	// Default FTP port
	host := parsedURL.Host
	if !strings.Contains(host, ":") {
		host = host + ":21"
	}

	// Connect to FTP server
	conn, err := ftp.Dial(host, ftp.DialWithTimeout(30*time.Second))
	if err != nil {
		return fmt.Errorf("failed to connect to FTP server: %w", err)
	}
	defer conn.Quit()

	// Login (use anonymous if no credentials provided)
	username := "anonymous"
	password := "anonymous"
	if parsedURL.User != nil {
		username = parsedURL.User.Username()
		if pass, hasPass := parsedURL.User.Password(); hasPass {
			password = pass
		}
	}

	if err := conn.Login(username, password); err != nil {
		return fmt.Errorf("failed to login to FTP server: %w", err)
	}

	// Retrieve the file
	resp, err := conn.Retr(parsedURL.Path)
	if err != nil {
		return fmt.Errorf("failed to retrieve file from FTP server: %w", err)
	}
	defer resp.Close()

	// Get file size if available
	var fileSize int64
	if entry, err := conn.GetEntry(parsedURL.Path); err == nil {
		fileSize = int64(entry.Size)
	}

	// Create the output file
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	// Create progress writer
	progress := NewProgressWriter(out, fileSize)

	// Write the response to file with progress tracking
	_, err = io.Copy(progress, resp)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	// Finish progress display
	progress.Finish()

	return nil
}
