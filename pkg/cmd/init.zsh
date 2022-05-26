zle -N qwy::complete
qwy::complete() {
    eval "$(\command qwy expand --lbuffer="$LBUFFER" --rbuffer="$RBUFFER")"
    zle reset-prompt
}
