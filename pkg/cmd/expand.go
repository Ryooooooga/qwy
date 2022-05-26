package cmd

import (
	"fmt"

	"github.com/Ryooooooga/qwy/pkg/command"
	"github.com/Ryooooooga/qwy/pkg/config"
	"github.com/Ryooooooga/qwy/pkg/matcher"
)

type ExpandCmd struct {
	LBuffer string `name:"lbuffer" short:"l" required:"" help:"$LBUFFER"`
	RBuffer string `name:"rbuffer" short:"r" required:"" help:"$RBUFFER"`
}

func (c *ExpandCmd) Run() error {
	config, err := config.LoadConfig()
	if err != nil {
		return err
	}

	completion, err := matcher.FindMatchedCompletion(config.Completions, c.LBuffer, c.RBuffer)
	if err != nil {
		return err
	}
	if completion == nil {
		return nil
	}

	lastArgIndex := command.LastArgumentIndex(c.LBuffer)
	cmd := c.LBuffer[:lastArgIndex]
	query := c.LBuffer[lastArgIndex:]

	fmt.Printf("expand: %#v\ncmd: %#v\nquery: %#v\n", completion, cmd, query)
	return nil
}
