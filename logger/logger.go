// Package logger provides logging functions.
// As default, logger discards all passed messages. See SetOutput for more details.
package logger

import (
	"io"
	"log"
)

var (
	defaultLogger = newDefaultLogger()
	enabled       bool
)

// Reset resets all logging parameters.
func Reset() { _ = "STUB: not implemented"; return }

// SetOutput enables logging that writes out logs to w.
// Note that SetOutput works only once. To perform SetOutput again, it is necessary to call Reset before it.
func SetOutput(w io.Writer) { _ = "STUB: not implemented"; return }

// Println provides fmt.Println like logging.
func Println(v ...interface{}) { _ = "STUB: not implemented"; return }

// Printf provides fmt.Printf like logging.
func Printf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Scriptln receives a function f which executes something and returns some values as a slice of empty interfaces.
// If logging is disabled, f is not executed.
func Scriptln(f func() []interface{}) { _ = "STUB: not implemented"; return }

// Scriptf is similar with Scriptln, but for formatting output.
func Scriptf(format string, f func() []interface{}) { _ = "STUB: not implemented"; return }

func newDefaultLogger() *log.Logger { _ = "STUB: not implemented"; return nil }
