package repl

import (
	"regexp"

	"github.com/ktr0731/evans/prompt"
)

var spaces = regexp.MustCompile(`\s+`)

type completer struct {
	cmds        map[string]commander
	completions map[string]func(args []string) (s []*prompt.Suggest)
}

// Complete completes suggestions from the input. In the completion, if an error is occurred, it will be ignored.
func (c *completer) Complete(d prompt.Document) (s []*prompt.Suggest) {
	_ = "STUB: not implemented"
	return nil
}

// TODO: We should consider about spaces used as a part of test.

// Ignore command name.

// return all commands if current input is first command name

// number of commands + help

func newCompleter(cmds map[string]commander) *completer { _ = "STUB: not implemented"; return nil }
