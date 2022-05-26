package cmd

import (
	_ "embed"
	"fmt"
)

type InitCmd struct {
}

//go:embed init.zsh
var initScript string

func (c *InitCmd) Run() error {
	fmt.Print(initScript)
	return nil
}
