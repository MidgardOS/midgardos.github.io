if [[ -f /usr/lib/grep/grepconf.sh ]]; then
    /usr/lib/grep/grepconf.sh -c || return
fi

alias xzgrep='xzgrep --color=auto' 2>/dev/null
alias xzegrep='xzegrep --color=auto' 2>/dev/null
alias xzfgrep='xzfgrep --color=auto' 2>/dev/null
