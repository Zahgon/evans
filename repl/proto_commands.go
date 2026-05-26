package repl

import (
	"io"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type descCommand struct{}

func (c *descCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }

func (c *descCommand) Help() string { _ = "STUB: not implemented"; return "" }

func (c *descCommand) FlagSet() (*pflag.FlagSet, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *descCommand) Validate(args []string) error { _ = "STUB: not implemented"; return nil }

func (c *descCommand) Run(w io.Writer, args []string) error { _ = "STUB: not implemented"; return nil }

func presentTypeName(f protoreflect.FieldDescriptor) string { _ = "STUB: not implemented"; return "" }
