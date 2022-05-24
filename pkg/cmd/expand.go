package cmd

import (
	"fmt"
)

type ExpandCmd struct {
	LBuffer string `name:"lbuffer" short:"l" required help:"$LBUFFER"`
	RBuffer string `name:"rbuffer" short:"r" required help:"$RBUFFER"`
}

func (c *ExpandCmd) Run() error {
	fmt.Println("expand", c)
	return nil
}
