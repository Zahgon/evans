// Package cui defines charcter user interfaces for I/O.
package cui

import (
	"io"
)

// UI provides formatted I/O interfaces.
// It is used from Evans's standard I/O and CLI mode I/O.
type UI interface {
	Output(s string)
	Info(s string)
	Warn(s string)
	Error(s string)

	Writer() io.Writer
}

// New creates a new UI with passed options.
func New(opts ...Option) UI {
	_ = "STUB: not implemented"
	// Creates a new UI with stdin, stdout, stderr.
	return *new(UI)
}

type basicUI struct {
	writer, errWriter io.Writer
}

// Output writes out the passed argument s to Writer with a line break.
func (u *basicUI) Output(s string) { _ = "STUB: not implemented"; return }

// Info is the same as Output, but distinguish these for composition.
func (u *basicUI) Info(s string) {
	_ = "STUB: not implemented"

	// Warn is the same as Output, but distinguish these for composition.
	return
}

func (u *basicUI) Warn(s string) {
	_ = "STUB: not implemented"

	// Error writes out the passed argument s to ErrWriter with a line break.
	return
}

func (u *basicUI) Error(s string) { _ = "STUB: not implemented"; return }

// Writer returns an io.Writer which is used in u.
func (u *basicUI) Writer() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

type coloredUI struct {
	UI
}

// NewColored wraps provided `ui` with coloredUI.
// If `ui` is *coloredUI, NewColored returns it as it is.
// Colored output works fine in Windows environment.
func NewColored(ui UI) UI { _ = "STUB: not implemented"; return *new(UI) }

// Info is the same as New, but colored.
func (u *coloredUI) Info(s string) { _ = "STUB: not implemented"; return }

// Warn is the same as New, but colored.
func (u *coloredUI) Warn(s string) { _ = "STUB: not implemented"; return }

// Error is the same as New, but colored.
func (u *coloredUI) Error(s string) { _ = "STUB: not implemented"; return }
