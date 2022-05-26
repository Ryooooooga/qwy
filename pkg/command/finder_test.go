package command_test

import (
	"strings"
	"testing"

	"github.com/Ryooooooga/qwy/pkg/command"
	"github.com/Ryooooooga/qwy/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestBuildFinderCommand(t *testing.T) {
	config, err := config.LoadConfigFromText(`
finder:
  --exit-0: true

completions:
  - description: empty

  - description: options
    finder:
      --exit-0: false
      --preview: "cat {}"
      +i: true
      -n: 2
      --multi: 4294967296
      --bind: ["ctrl-d:print-query", "ctrl-p:replace-query"]
`)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, 2, len(config.Completions))

	actual, err := command.BuildFinderCommand(config.FinderCommand, `""`, config.Completions[0].Finder)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, `fzf --query "" --exit-0`, actual)

	actual, err = command.BuildFinderCommand(config.FinderCommand, "x", config.Completions[1].Finder)
	if err != nil {
		t.Error(err)
		return
	}
	assert.True(t, strings.HasPrefix(actual, "fzf "))
	assert.Contains(t, actual, " --query x")
	assert.NotContains(t, actual, " --exit-0")
	assert.Contains(t, actual, " --preview 'cat {}'")
	assert.Contains(t, actual, " +i")
	assert.Contains(t, actual, " -n 2")
	assert.Contains(t, actual, " --multi 4294967296")
	assert.Contains(t, actual, " --bind ctrl-d:print-query --bind ctrl-p:replace-query")
}
