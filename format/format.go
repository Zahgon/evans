// Package format provides formatting APIs for display processed result application did.
package format

import (
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// ResponseFormatter provides formatting feature for gRPC response.
type ResponseFormatter struct {
	enrich bool

	impl ResponseFormatterInterface
}

func (f *ResponseFormatter) Format(s *status.Status, header, trailer metadata.MD, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *ResponseFormatter) FormatHeader(header metadata.MD) { _ = "STUB: not implemented"; return }

func (f *ResponseFormatter) FormatMessage(v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *ResponseFormatter) FormatTrailer(status *status.Status, trailer metadata.MD) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *ResponseFormatter) Done() error { _ = "STUB: not implemented"; return nil }

// NewResponseFormatter formats gRPC response with a specific formatter.
// If enrich is false, the formatter prints only messages.
// Or else, it prints all includes headers, messages, trailers and status.
func NewResponseFormatter(f ResponseFormatterInterface, enrich bool) *ResponseFormatter {
	_ = "STUB: not implemented"
	return nil
}

// ResponseFormatterInterface is an interface for formatting gRPC response.
type ResponseFormatterInterface interface {
	// FormatHeader formats the response header.
	FormatHeader(header metadata.MD)
	// FormatMessage formats the response message (body).
	FormatMessage(v interface{}) error
	// FormatStatus formats the response status.
	FormatStatus(status *status.Status) error
	// FormatTrailer formats the response trailer.
	FormatTrailer(trailer metadata.MD)
	// Done indicates all response information is formatted.
	// The client of ResponseFormatter should call it at the end.
	Done() error
}
