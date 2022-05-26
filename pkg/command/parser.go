package command

type state int

const (
	stateSpace state = iota + 1
	stateArg
	stateEscaped
)

func isSpace(c rune) bool {
	switch c {
	case '\t', '\n', ' ':
		return true
	default:
		return false
	}
}

func LastArgumentIndex(command string) int {
	index := 0
	state := stateSpace
	for i, c := range command {
		switch state {
		case stateSpace:
			switch {
			case isSpace(c):
			case c == '\\':
				index = i
				state = stateEscaped
			default:
				index = i
				state = stateArg
			}
		case stateArg:
			switch {
			case isSpace(c):
				state = stateSpace
			case c == '\\':
				state = stateEscaped
			default:
			}
		case stateEscaped:
		}
	}
	if state == stateSpace {
		index = len(command)
	}
	return index
}
