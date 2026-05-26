// Package prompt provides the prompt interface and its implementation.
package prompt

import (
	"github.com/chzyer/readline"
	goprompt "github.com/ktr0731/go-prompt"
	"github.com/pkg/errors"
)

type stdout struct{}

func (s *stdout) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *stdout) Close() error { _ = "STUB: not implemented"; return nil }

func init() {
	// Override readline.Stdout to suppress ringing bell.
	// See more details: https://github.com/manifoldco/promptui/issues/49#issuecomment-428801411
	readline.Stdout = &stdout{}
}

// InitialColor is the initial color for a prompt prefix.
var (
	ColorInitial = Color(goprompt.DarkGreen)
	ColorBlue    = Color(goprompt.Blue)
)

var (
	ErrSkip  = errors.New("skip")
	ErrAbort = errors.New("abort")
)

// Color represents a valid color for a prompt prefix.
type Color goprompt.Color

// Next returns the next color of c. Note that Next will circular if c is the end of colors.
func (c *Color) Next() {
	_ = "STUB: not implemented"

	// NextVal is the same as Next, but return color as value.
	return
}

func (c *Color) NextVal() Color { _ = "STUB: not implemented"; return *new(Color) }

type Prompt interface {
	// Input reads keyboard input.
	// If ctrl+d is entered, Input returns io.EOF.
	// If ctrl+c is entered, Input returns ErrAbort.
	Input() (string, error)
	Select(message string, options []string) (idx int, selected string, _ error)

	// SetPrefix changes the current prefix to the passed one.
	SetPrefix(prefix string)

	// SetPrefixColor changes the current color to the passed one.
	SetPrefixColor(color Color)

	// SetCompleter set a completer for prompt completion.
	SetCompleter(c Completer)

	// GetCommandHistory gets a command history. The order of history is asc.
	GetCommandHistory() []string
}

// New instantiates a new Prompt implementation. New will be replaced when e2egen command is executed.
// Initially, Prompt doesn't have a prefix, so you have to call SetPrefix for displaying it.
var New = newPrompt

func newPrompt(opts ...Option) Prompt { _ = "STUB: not implemented"; return *new(Prompt) }

type prompt struct {
	prefix         string
	prefixColor    Color
	completer      Completer
	commandHistory []string
	options        []goprompt.Option

	// Treat prompt functions as fields for testing.
	InputFunc  func(prefix string, completer goprompt.Completer, opts ...goprompt.Option) (string, error)
	SelectFunc func(message string, options []string) (int, string, error)
}

func (p *prompt) Input() (in string, err error) { _ = "STUB: not implemented"; return "", nil }

func (p *prompt) Select(message string, options []string) (int, string, error) {
	_ = "STUB: not implemented"
	return 0, "", nil
}

func (p *prompt) SetPrefix(prefix string) { _ = "STUB: not implemented"; return }

func (p *prompt) SetPrefixColor(color Color) { _ = "STUB: not implemented"; return }

func (p *prompt) SetCompleter(c Completer) { _ = "STUB: not implemented"; return }

func (p *prompt) GetCommandHistory() []string { _ = "STUB: not implemented"; return nil }

func (p *prompt) livePrefix() (string, bool) {
	_ = "STUB: not implemented"
	return "",

		// Completer is a mechanism that provides REPL completion.
		false
}

type Completer interface {
	// Complete receives d that is a piece of input, and returns some suggestions.
	// The prompt shows suggestions from it.
	Complete(d Document) []*Suggest
}

// Document is a piece of input that has several information such that
// text before the cursor, a work before the cursor, etc.
type Document interface {
	GetWordBeforeCursor() string
	TextBeforeCursor() string
}

// Suggest is a suggestion for the completion.
type Suggest struct {
	goprompt.Suggest
}

// NewSuggestion returns a new *Suggest from text and description.
func NewSuggestion(text, description string) *Suggest { _ = "STUB: not implemented"; return nil }

// FilterHasPrefix filters s by whether have sub as the prefix.
// If ignoreCase is true, differences between upper and lower casing are ignored.
func FilterHasPrefix(s []*Suggest, sub string, ignoreCase bool) []*Suggest {
	_ = "STUB: not implemented"
	return nil
}

func toGoPromptCompleter(c Completer) goprompt.Completer {
	_ = "STUB: not implemented"
	return *new(goprompt.Completer)
}

func fromGoPromptSuggestions(s []goprompt.Suggest) []*Suggest {
	_ = "STUB: not implemented"
	return nil
}

func fromPromptSuggestions(s []*Suggest) []goprompt.Suggest { _ = "STUB: not implemented"; return nil }
