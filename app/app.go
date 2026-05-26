// Package app provides the entrypoint for Evans.
package app

import (
	"io"

	"github.com/ktr0731/evans/config"
	"github.com/ktr0731/evans/cui"
	"github.com/spf13/pflag"
)

// App is the root component for running the application.
type App struct {
	cui cui.UI
	cmd *command
}

// New instantiates a new App instance. ui must not be a nil.
// Note that cui is also used for the REPL UI if the mode is REPL mode.
func New(ui cui.UI) *App { _ = "STUB: not implemented"; return nil }

// Run starts the application. The return value means the exit code.
func (a *App) Run(args []string) int {
	_ = "STUB: not implemented"
	// Currently, Evans is migrating to new-style command-line interface.
	// So, there are both of old-style and new-style command-line interfaces in this version.
	return 0
}

// Hack.

// Sub commands for new-style interface.
// If an arg named "cli" or "repl" is passed, it is regarded as a sub-command of new-style.

// If the help flags is passed, call registerNewCommands for display sub-command helps.

// printUsage shows the command usage text to cui.Writer and exit. Do not call it before calling parseFlags.
func printUsage(cmd interface{ Help() error }) {
	_ = "STUB: not implemented"
	// Help never return errors.
	return
}

func printVersion(w io.Writer) { _ = "STUB: not implemented"; return }

// mergedConfig represents the conclusive config. Common config items are stored to *config.Config.
// Flags that can be specified by command line only are represented as fields.
type mergedConfig struct {
	*config.Config

	// The RPC name that want to call.
	call string
	// The input that is used for CLI mode. Empty if the input is stdin.
	file string
	// Explicit using CLI mode.
	cli bool

	// Explicit using REPL mode.
	repl bool
}

func mergeConfig(fs *pflag.FlagSet, flags *flags, protos []string) (*mergedConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
