// Package json provides a JSON presenter that formatting.
package json

// Presenter is a presenter that formats v into JSON string.
type Presenter struct {
	indent string
}

// Format formats v into JSON string.
func (p *Presenter) Format(v interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// NewPresenter instantiates a JSON presenter.
// If indent is not empty, Format indents the output.
func NewPresenter(indent string) *Presenter { _ = "STUB: not implemented"; return nil }
