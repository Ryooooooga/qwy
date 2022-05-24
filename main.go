package main

import (
	"fmt"

	"github.com/Ryooooooga/qwy/pkg/cmd"
	"github.com/alecthomas/kong"
)

const (
	QWY string = "qwy"
)

var (
	version string = "dev"
	commit  string = "HEAD"
	date    string = "unknown"
)

type VersionFlag bool

func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Printf("%s version %s (rev: %s) built at %s\n", QWY, version, commit, date)
	app.Exit(0)
	return nil
}

var cli struct {
	Init    cmd.InitCmd   `cmd:"" help:"Used to setup the plugin"`
	Expand  cmd.ExpandCmd `cmd:"" help:"Expand completions"`
	Version VersionFlag   `short:"v" help:"Print the version"`
}

func main() {
	ctx := kong.Parse(
		&cli,
		kong.Name(QWY),
		kong.Description("Fuzzy ZSH completion plugin"),
		kong.UsageOnError(),
	)

	err := ctx.Run(&kong.Context{})
	ctx.FatalIfErrorf(err)
}
