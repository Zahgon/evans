package app

import (
	"io"

	"github.com/ktr0731/evans/cui"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type command struct {
	*cobra.Command

	flags *flags
	ui    cui.UI
}

// registerNewCommands registers sub-commands for new-style interface.
func (c *command) registerNewCommands() { _ = "STUB: not implemented"; return }

// runFunc is a common entrypoint for Run func.
func runFunc(
	flags *flags,
	f func(cmd *cobra.Command, cfg *mergedConfig) error,
) func(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Help is processed by cobra.

// For backward-compatibility.

// Pass Flags instead of LocalFlags because the config is merged with common and local flags.

// The entrypoint for the command.

func newOldCommand(flags *flags, ui cui.UI) *command { _ = "STUB: not implemented"; return nil }

func bindFlags(f *pflag.FlagSet, flags *flags, w io.Writer) { _ = "STUB: not implemented"; return }

// Flags used by old-style only.
// Hidden is enabled only the root command (see printOptions).

func newCLICommand(flags *flags, ui cui.UI) *cobra.Command { _ = "STUB: not implemented"; return nil }

// For backward-compatibility.
// If the method is specified by passing --call option, use it.

func newREPLCommand(flags *flags, ui cui.UI) *cobra.Command { _ = "STUB: not implemented"; return nil }

func runREPLCommand(cfg *mergedConfig, ui cui.UI) error { _ = "STUB: not implemented"; return nil }

// Run update checker asynchronously.

// Always call cancel() because it is hope to abort update checking if REPL mode is finished
// before update checking. If the update check is finished before REPL mode, cancel does nothing.

func initFlagSet(f *pflag.FlagSet, w io.Writer) { _ = "STUB: not implemented"; return }

func printOptions(w io.Writer, cmd *cobra.Command, inheritedFlags []string) {
	_ = "STUB: not implemented"
	return
}

// Always show --help text.

// usage is the generator for usage output.
func usageFunc(out io.Writer, inheritedFlags []string) func(*cobra.Command, []string) {
	_ = "STUB: not implemented"
	return nil
}

// Ignore help command.
