// Package repl provides a REPL environment for REPL mode.
package repl

import (
	"context"

	"github.com/ktr0731/evans/config"
	"github.com/ktr0731/evans/cui"
	"github.com/ktr0731/evans/prompt"
)

// REPL represents a REPL mechanism.
type REPL struct {
	cfg       *config.REPL
	serverCfg *config.Server
	prompt    prompt.Prompt
	ui        cui.UI

	cmds    map[string]commander
	aliases map[string]string
}

var commands = map[string]commander{
	"call":    &callCommand{},
	"service": &serviceCommand{},
	"header":  &headerCommand{},
	"package": &packageCommand{},
	"show":    &showCommand{},
	"exit":    &exitCommand{},

	// Depends to Protocol Buffers.
	"desc": &descCommand{},
}

// New instantiates a new REPL instance. New always calls p.SetPrefix for display the server addr.
// New may return an error if some of passed arguments are invalid.
func New(cfg *config.Config, p prompt.Prompt, ui cui.UI, pkgName, svcName string) (*REPL, error) {
	_ = "STUB: not implemented"

	// Each value must be a key of cmds.
	return nil, nil
}

// Run starts the read-eval-print-loop.
func (r *REPL) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Break line.

func (r *REPL) runCommand(cmdName string, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check whether cmdName is an alias for a command.

func (r *REPL) printSplash(p string) { _ = "STUB: not implemented"; return }

func (r *REPL) makePrefix() string { _ = "STUB: not implemented"; return "" }

func (r *REPL) helpText() string {
	_ = "STUB: not implemented"

	// slice of [name, synopsis]
	return ""
}
