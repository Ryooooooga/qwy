package matcher

import (
	"regexp"

	"github.com/Ryooooooga/qwy/pkg/config"
)

type Captures map[string]string

func FindMatchedCompletion(completions config.Completions, lbuffer, rbuffer string) (*config.Completion, Captures, error) {
	for _, c := range completions {
		matched, captures, err := isMatched(c, lbuffer, rbuffer)
		if err != nil {
			return nil, Captures{}, err
		}
		if matched {
			return c, captures, nil
		}
	}

	return nil, Captures{}, nil
}

func isMatched(completion *config.Completion, lbuffer, rbuffer string) (bool, Captures, error) {
	for _, pattern := range completion.Patterns {
		r, err := regexp.Compile(pattern)
		if err != nil {
			return false, Captures{}, err
		}

		matches := r.FindStringSubmatch(lbuffer)
		if len(matches) == 0 {
			continue
		}

		captures := Captures{}
		for i, name := range r.SubexpNames() {
			if len(name) > 0 {
				captures[name] = matches[i]
			}
		}
		return true, captures, nil
	}

	return false, Captures{}, nil
}
