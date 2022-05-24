package cmd

import (
	"fmt"
)

type InitCmd struct {
}

func (c *InitCmd) Run() error {
	fmt.Println("init")
	return nil
}
