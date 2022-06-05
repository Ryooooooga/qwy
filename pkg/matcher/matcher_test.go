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

  - description: cd
    patterns:
      - ^cd\s+(?P<cap>.*)
`)
	assert.Nil(t, err)

	scenarios := []struct {
		lbuffer             string
		rbuffer             string
		expectedDescription string
		expectedCaptures    matcher.Captures
	}{
		{
			lbuffer:             "git ",
			rbuffer:             "",
			expectedDescription: "git",
			expectedCaptures:    matcher.Captures{},
		},
		{
			lbuffer:             "git add ",
			rbuffer:             "",
			expectedDescription: "git-add",
			expectedCaptures:    matcher.Captures{},
		},
		{
			lbuffer:             "cd ~/",
			rbuffer:             "",
			expectedDescription: "cd",
			expectedCaptures:    matcher.Captures{"cap": "~/"},
		},
		{
			lbuffer:             "not-matched",
			rbuffer:             "",
			expectedDescription: "",
			expectedCaptures:    matcher.Captures{},
		},
	}

	for i, s := range scenarios {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			c, cap, err := matcher.FindMatchedCompletion(config.Completions, s.lbuffer, s.rbuffer)
			assert.Nil(t, err)

			if len(s.expectedDescription) > 0 {
				assert.NotNil(t, c)
				assert.Equal(t, s.expectedDescription, c.Description)
				assert.Equal(t, s.expectedCaptures, cap)
			} else {
				assert.Nil(t, c)
			}
		})
	}
}
