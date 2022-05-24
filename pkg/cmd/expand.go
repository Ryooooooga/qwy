package cmd

import (
	"fmt"

	"github.com/Ryooooooga/qwy/pkg/config"
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

	fmt.Println("expand", c, config)
	return nil
}
