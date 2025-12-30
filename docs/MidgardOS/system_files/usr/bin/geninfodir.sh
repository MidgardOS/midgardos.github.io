#!/usr/bin/env bash

# This script generates an infodir file for all info files on the system.
# It should be run with root privileges to ensure all info files are accessible.

INFO_DIR="/usr/share/info"
INFODIR_FILE="$INFO_DIR/dir"

# Check if the info directory exists
if [ ! -d "$INFO_DIR" ]; then
    echo "Info directory $INFO_DIR does not exist."
    exit 1
fi

# Create or clear the infodir file
: > "$INFODIR_FILE"
echo "Generating infodir file at $INFODIR_FILE..."
# Find all .info files and generate the infodir entries
find "$INFO_DIR" -name "*.info" | while read -r infofile; do
    # Extract the base name and title from the info file
    basename=$(basename "$infofile" .info)
    title=$(grep -m 1 '^@title' "$infofile" | sed 's/^@title //')
    if [ -n "$title" ]; then
        echo "$basename\t$title" >> "$INFODIR_FILE"
    fi
done
echo "Infodir file generated successfully."
echo "You may need to run 'install-info $INFODIR_FILE $INFO_DIR/dir' to update the system info directory."

pushd "$INFO_DIR"
    rm -v dir
    for f in *; do
        install-info $f dir 2>/dev/null
    done
popd

exit 0
