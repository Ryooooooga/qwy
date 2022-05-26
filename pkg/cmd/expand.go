package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/Ryooooooga/qwy/pkg/command"
	"github.com/Ryooooooga/qwy/pkg/config"
	"github.com/Ryooooooga/qwy/pkg/matcher"
	"gopkg.in/alessio/shellescape.v1"
)

const (
	LBUFFER  string = "LBUFFER"
	RBUFFER  string = "RBUFFER"
	PREFIX   string = "__qwy_prefix"
	QUERY    string = "__qwy_query"
	SOURCE   string = "__qwy_source"
	FINDER   string = "__qwy_finder"
	CALLBACK string = "__qwy_callback"
	OUTPUT   string = "__qwy_output"
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
	prefix := c.LBuffer[:lastArgIndex]
	query := c.LBuffer[lastArgIndex:]

	escapedQuery := fmt.Sprintf(`"${%s}"`, QUERY)

	finderCommand, err := command.BuildFinderCommand(config.FinderCommand, escapedQuery, completion.Finder)
	if err != nil {
		return err
	}

	writeScript(os.Stdout, prefix, query, completion.Source, finderCommand, completion.Callback)
	return nil
}

func writeScript(w io.Writer, prefix, query, sourceCommand, finderCommand, callbackCommand string) {
	fmt.Fprintf(w, "local %s=%s;\n", PREFIX, shellescape.Quote(prefix))
	fmt.Fprintf(w, "local %s=%s;\n", QUERY, shellescape.Quote(query))
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
