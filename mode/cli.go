package mode

import (
	"context"
	"io"
	"os"

	"github.com/ktr0731/evans/config"
	"github.com/ktr0731/evans/cui"
)

// DefaultCLIReader is the reader that is read for inputting request values. It is exported for E2E testing.
var DefaultCLIReader io.Reader = os.Stdin

// CLIInvoker represents an invokable function for CLI mode.
type CLIInvoker func(context.Context) error

type CallCLIInvokerOption struct {
	Headers      config.Header
	Enrich       bool
	EmitDefaults bool
	FilePath     string // If empty, the invoker tries to read input from stdin.
	FormatType   string
}

// NewCallCLIInvoker returns an CLIInvoker implementation for calling RPCs.
func NewCallCLIInvoker(ui cui.UI, methodName string, opt *CallCLIInvokerOption) (CLIInvoker, error) {
	_ = "STUB: not implemented"
	return *new(CLIInvoker), nil
}

// Try to parse methodName as a fully-qualified method name.
// If it is valid, use its fully-qualified service.

func NewListCLIInvoker(ui cui.UI, fqn, format string) CLIInvoker {
	_ = "STUB: not implemented"
	return *new(CLIInvoker)
}

// A fully-qualified method name is passed.
// Return as it is (same behavior as grpc_cli).

// Parse as a fully-qualified service name.

// Return commonErr because UsePackage will be deprecated.

func NewDescribeCLIInvoker(ui cui.UI, fqn string) CLIInvoker {
	_ = "STUB: not implemented"
	return *new(CLIInvoker)
}

// RunAsCLIMode starts Evans as CLI mode.
func RunAsCLIMode(cfg *config.Config, invoker CLIInvoker) error {
	_ = "STUB: not implemented"
	return nil
}

// IsCLIMode returns whether Evans is launched as CLI mode or not.
func IsCLIMode(file string) bool { _ = "STUB: not implemented"; return false }

func isFullyQualifiedMethodName(s string) bool { _ = "STUB: not implemented"; return false }
