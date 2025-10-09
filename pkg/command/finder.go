package command

import (
	"fmt"
	"strings"

	"github.com/Ryooooooga/qwy/pkg/config"
	"github.com/pkg/errors"
	"gopkg.in/alessio/shellescape.v1"
)

func BuildFinderCommand(finder string, finderOptions config.FinderOptions) (string, error) {
	var b strings.Builder

	b.WriteString(finder)

	for name, value := range finderOptions {
		switch v := value.(type) {
		case []any:
			if err := writeArrayOption(&b, name, v); err != nil {
				return "", err
			}
		default:
			if err := writePrimitiveOption(&b, name, v, name); err != nil {
				return "", err
			}
		}
	}

	return b.String(), nil
}

func writeArrayOption(b *strings.Builder, name string, values []any) error {
	for i, value := range values {
		if err := writePrimitiveOption(b, name, value, fmt.Sprintf("%s[%d]", name, i)); err != nil {
			return err
		}
	}
	return nil
}

func writePrimitiveOption(b *strings.Builder, name string, value any, selector string) error {
	b.WriteString(" ")

	switch v := value.(type) {
	case bool:
		if v {
			b.WriteString(shellescape.Quote(name))
		}
	case string:
		_, _ = fmt.Fprintf(b, "%s %s", shellescape.Quote(name), quoteValue(v))
	case int:
		_, _ = fmt.Fprintf(b, "%s %d", shellescape.Quote(name), v)
	default:
		return errors.Errorf("finder.%s must be string, int, or bool: %v", selector, value)
	}

	return nil
}

var quoteReplacer *strings.Replacer = strings.NewReplacer(
	`\$`, "\\$",
	"\\`", "\\\\`",
	`\`, `\\`,
	`"`, `\"`,
)

func quoteValue(s string) string {
	return `"` + quoteReplacer.Replace(s) + `"`
}
