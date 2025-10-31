if [[ -f /usr/lib/grep/grepconf.sh ]]; then
    /usr/lib/grep/grepconf.sh -c || return
fi
alias zgrep='zgrep --color=auto' 2>/dev/null
alias zfgrep='zfgrep --color=auto' 2>/dev/null
alias zegrep='zegrep --color=auto' 2>/dev/null
