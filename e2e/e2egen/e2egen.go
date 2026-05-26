// Package main provides a generator for E2E test cases.
//
// e2egen generates a new testcase stub from execution result. e2egen accepts following args.
//
//	e2egen <file> <Evans flags>
//
// file is the source code of REPL E2E test. Updated code is written to it. Evans flags are options of Evans.
//
// First, e2egen tries to read source code from file.
// Second, e2egen launches Evans with Evans flags. In this time, e2egen records all input from the prompt internally.
// Third, e2egen generates updated source code that appended the new testcase.
package main

import (
	"io"
	"os"

	"github.com/ktr0731/evans/prompt"
)

func main() {
	os.Exit(realMain())
}

func realMain() int { _ = "STUB: not implemented"; return 0 }

func generateFile(w io.Writer, src, testCaseName string, input, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Add a large value.

type recorderPrompt struct {
	prompt.Prompt

	inputHistory []string
}

func (p *recorderPrompt) Input() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (p *recorderPrompt) Select(message string, options []string) (int, string, error) {
	_ = "STUB: not implemented"
	return 0, "", nil
}

func filterArgs(args []string) []string { _ = "STUB: not implemented"; return nil }

// Skip arg of --port and itself.
