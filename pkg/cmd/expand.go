package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/Ryooooooga/qwy/pkg/command"
	"github.com/Ryooooooga/qwy/pkg/config"
	"github.com/Ryooooooga/qwy/pkg/maps"
	"github.com/Ryooooooga/qwy/pkg/matcher"
	"gopkg.in/alessio/shellescape.v1"
)

const (
	LBUFFER  string = "LBUFFER"
	RBUFFER  string = "RBUFFER"
	QUERY    string = "query"
	PREFIX   string = "__qwy_prefix"
	SOURCE   string = "__qwy_source"
	FINDER   string = "__qwy_finder"
	CALLBACK string = "__qwy_callback"
	OUTPUT   string = "__qwy_output"
)

var defaultFinderOptions config.FinderOptions = config.FinderOptions{
	"--query": `${(Q)query}`,
}

type ExpandCmd struct {
	LBuffer string `name:"lbuffer" short:"l" required:"" help:"$LBUFFER"`
	RBuffer string `name:"rbuffer" short:"r" required:"" help:"$RBUFFER"`
}

func (c *ExpandCmd) Run() error {
	config, err := config.LoadConfig()
	if err != nil {
		return err
	}

	completion, captures, err := matcher.FindMatchedCompletion(config.Completions, c.LBuffer, c.RBuffer)
	if err != nil {
		return err
	}
	if completion == nil {
		return nil
	}

	lastArgIndex := command.LastArgumentIndex(c.LBuffer)
	prefix := c.LBuffer[:lastArgIndex]
	query := c.LBuffer[lastArgIndex:]

	if _, ok := captures[QUERY]; !ok {
		captures[QUERY] = query
	}

	finderOptions := maps.Merge(defaultFinderOptions, config.Finder, completion.Finder)
	finderCommand, err := command.BuildFinderCommand(config.FinderCommand, finderOptions)
	if err != nil {
		return err
	}

	writeScript(os.Stdout, prefix, captures, completion.Source, finderCommand, completion.Callback)
	return nil
}

func writeScript(w io.Writer, prefix string, captures matcher.Captures, sourceCommand, finderCommand, callbackCommand string) {
	fmt.Fprintf(w, "local %s;\n", OUTPUT)

	for name, capture := range captures {
		fmt.Fprintf(w, "local %s=%s;\n", name, shellescape.Quote(capture))
	}

	fmt.Fprintf(w, "local %s=%s;\n", PREFIX, shellescape.Quote(prefix))
	fmt.Fprintf(w, "local %s=%s;\n", SOURCE, shellescape.Quote(sourceCommand))
	fmt.Fprintf(w, "local %s=%s;\n", FINDER, shellescape.Quote(finderCommand))
	if len(callbackCommand) > 0 {
		fmt.Fprintf(w, "local %s=%s;\n", CALLBACK, shellescape.Quote(callbackCommand))
	}

	fmt.Fprintf(w, `if %s="$(\builtin eval "${%s}" | \builtin eval "${%s}")";then `, OUTPUT, SOURCE, FINDER)
	if len(callbackCommand) > 0 {
		fmt.Fprintf(w, `%s="$(\builtin eval "${%s}" <<<"${%s}")";`, OUTPUT, CALLBACK, OUTPUT)
	}
	fmt.Fprintf(w, `%s="${%s}${%s}";`, LBUFFER, PREFIX, OUTPUT)
	fmt.Fprintln(w, "fi")
}
