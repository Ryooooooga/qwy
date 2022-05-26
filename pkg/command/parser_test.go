package command_test

import (
	"testing"

	"github.com/Ryooooooga/qwy/pkg/command"
	"github.com/stretchr/testify/assert"
)

func TestLastArgumentIndex(t *testing.T) {
	scenarios := []struct {
		input    string
		expected int
	}{
		{
			input:    "",
			expected: 0,
		},
		{
			input:    "  ",
			expected: 2,
		},
		{
			input:    "echo",
			expected: 0,
		},
		{
			input:    "\techo",
			expected: 1,
		},
		{
			input:    "echo ",
			expected: 5,
		},
		{
			input:    "echo a",
			expected: 5,
		},
		{
			input:    "echo a ",
			expected: 7,
		},
		{
			input:    `echo a\ b`,
			expected: 5,
		},
	}

	for _, s := range scenarios {
		t.Run(s.input, func(t *testing.T) {
			actual := command.LastArgumentIndex(s.input)
			assert.Equal(t, s.expected, actual)
		})
	}
}
