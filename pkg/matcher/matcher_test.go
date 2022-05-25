package matcher_test

import (
	"fmt"
	"testing"

	"github.com/Ryooooooga/qwy/pkg/config"
	"github.com/Ryooooooga/qwy/pkg/matcher"
	"github.com/stretchr/testify/assert"
)

func TestFindMatchedCompletion(t *testing.T) {
	config, err := config.LoadConfigFromText(`
completions:
  - description: empty

  - description: git-add
    patterns:
      - ^git\s+add\s

  - description: git
    patterns:
      - ^git\s
`)
	assert.Nil(t, err)

	scenarios := []struct {
		lbuffer  string
		rbuffer  string
		expected string
	}{
		{
			lbuffer:  "git ",
			rbuffer:  "",
			expected: "git",
		},
		{
			lbuffer:  "git add ",
			rbuffer:  "",
			expected: "git-add",
		},
		{
			lbuffer:  "not-matched",
			rbuffer:  "",
			expected: "",
		},
	}

	for i, s := range scenarios {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			c, err := matcher.FindMatchedCompletion(config.Completions, s.lbuffer, s.rbuffer)
			assert.Nil(t, err)

			if len(s.expected) > 0 {
				assert.NotNil(t, c)
				assert.Equal(t, s.expected, c.Description)
			} else {
				assert.Nil(t, c)
			}
		})
	}
}
