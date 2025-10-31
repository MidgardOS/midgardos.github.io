# color-grep initialization

if [[ -f /usr/lib/grep/grepconf.sh ]]; then
    /usr/lib/grep/grepconf.sh -c || return
fi

alias grep='grep --color=auto' 2>/dev/null
alias egrep='grep -E --color=auto' 2>/dev/null
alias fgrep='grep -F --color=auto' 2>/dev/null
