package matcher

import (
	"regexp"

	"github.com/Ryooooooga/qwy/pkg/config"
)

func FindMatchedCompletion(completions config.Completions, lbuffer, rbuffer string) (*config.Completion, error) {
	for _, c := range completions {
		matched, err := isMatched(c, lbuffer, rbuffer)
		if err != nil {
			return nil, err
		}
		if matched {
			return c, nil
		}
	}

	return nil, nil
}

func isMatched(completion *config.Completion, lbuffer, rbuffer string) (bool, error) {
	for _, pattern := range completion.Patterns {
		r, err := regexp.Compile(pattern)
		if err != nil {
			return false, err
		}
		if r.MatchString(lbuffer) {
			return true, nil
		}
	}

	return false, nil
}
