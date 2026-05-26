package repl

import (
	"io"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

var (
	errArgumentRequired = errors.New("argument required")
)

type commander interface {
	// Help returns a short help message.
	Help() string

	// Synopsis returns the usage of the command.
	Synopsis() string

	// FlagSet returns a flagset.
	// If the command doesn't have a flagset, the second returned value is false.
	// Note that FlagSet is for read-only, so don't modify it.
	FlagSet() (*pflag.FlagSet, bool)

	// Valdiate validates whether args satisfies preconditions for running the command.
	Validate(args []string) error

	// Run runs the command. The commander implementation writes something to w.
	// Caller must check no errors by calling Validate before call Run.
	Run(w io.Writer, args []string) error
}

type packageCommand struct{}

func (c *packageCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *packageCommand) Help() string { _ = "STUB: not implemented"; return "" }

func (c *packageCommand) FlagSet() (*pflag.FlagSet, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *packageCommand) Validate(args []string) error { _ = "STUB: not implemented"; return nil }

func (c *packageCommand) Run(_ io.Writer, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

type serviceCommand struct{}

func (c *serviceCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *serviceCommand) Help() string { _ = "STUB: not implemented"; return "" }

func (c *serviceCommand) FlagSet() (*pflag.FlagSet, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *serviceCommand) Validate(args []string) error { _ = "STUB: not implemented"; return nil }

func (c *serviceCommand) Run(_ io.Writer, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

type showCommand struct{}

func (c *showCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *showCommand) Help() string { _ = "STUB: not implemented"; return "" }

func (c *showCommand) FlagSet() (*pflag.FlagSet, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *showCommand) Validate(args []string) error { _ = "STUB: not implemented"; return nil }

func (c *showCommand) Run(w io.Writer, args []string) error { _ = "STUB: not implemented"; return nil }

type callCommand struct {
	enrich, digManually, bytesAsBase64, bytesAsQuotedLiterals, bytesFromFile, emitDefaults, repeatCall, addRepeatedManually bool
}

func (c *callCommand) FlagSet() (*pflag.FlagSet, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Disable help output when an error occurred.

func (c *callCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *callCommand) Help() string { _ = "STUB: not implemented"; return "" }

func (c *callCommand) Validate(args []string) error { _ = "STUB: not implemented"; return nil }

func (c *callCommand) Run(w io.Writer, args []string) error { _ = "STUB: not implemented"; return nil }

// Ensure only one of bytesAsBase64, bytesAsQuotedLiterals and bytesFromFile are not both set
// pflag doesn't suppport mutually exclusive flags (https://github.com/spf13/pflag/issues/270)

// here we create the request context
// we also add the call command flags here

type headerCommand struct {
	raw bool
}

func (c *headerCommand) FlagSet() (*pflag.FlagSet, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Disable help output when an error occurred.

func (c *headerCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *headerCommand) Help() string { _ = "STUB: not implemented"; return "" }

func (c *headerCommand) Validate(args []string) error { _ = "STUB: not implemented"; return nil }

func (c *headerCommand) Run(_ io.Writer, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove the key.

type exitCommand struct{}

func (c *exitCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *exitCommand) Help() string { _ = "STUB: not implemented"; return "" }

func (c *exitCommand) FlagSet() (*pflag.FlagSet, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *exitCommand) Validate([]string) error { _ = "STUB: not implemented"; return nil }

func (c *exitCommand) Run(io.Writer, []string) error { _ = "STUB: not implemented"; return nil }
