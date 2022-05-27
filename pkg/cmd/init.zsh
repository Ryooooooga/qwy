QWY_TRIGGER_KEY="${QWY_TRIGGER_KEY:-^I}"
QWY_DEFAULT_ACTION="${QWY_DEFAULT_ACTION-${$(\builtin bindkey "$QWY_TRIGGER_KEY")[(s: :w)2]}}"

\builtin zle -N qwy::complete
qwy::complete() {
    local out="$(\command qwy expand --lbuffer="${LBUFFER}" --rbuffer="${RBUFFER}")"
    if [[ -n "${out}" ]]; then
        \builtin eval "${out}"
        \builtin zle reset-prompt
    elif [[ -n "${QWY_DEFAULT_ACTION}" && "${QWY_DEFAULT_ACTION}" != "undefined-key" ]]; then
        \builtin zle "${QWY_DEFAULT_ACTION}"
    fi
}

\builtin bindkey "${QWY_TRIGGER_KEY}" qwy::complete
