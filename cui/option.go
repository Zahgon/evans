package cui

import (
	"io"
)

// Option represents an option for New.
type Option func(*basicUI)

// Writer modifies the default writer to w.
func Writer(w io.Writer) Option { _ = "STUB: not implemented"; return *new(Option) }

// ErrWriter modifies the default error writer to w.
func ErrWriter(ew io.Writer) Option { _ = "STUB: not implemented"; return *new(Option) }
