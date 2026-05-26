// Package json provides a JSON formatter implementation.
package json

import (
	"io"

	"github.com/golang/protobuf/jsonpb" //nolint:staticcheck
	"github.com/golang/protobuf/proto"  //nolint:staticcheck
	"github.com/ktr0731/evans/format"
	"github.com/ktr0731/evans/present"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails" // For calling RegisterType.
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// responseFormatter is a formatter that formats *usecase.GRPCResponse into a JSON object.
type responseFormatter struct {
	w io.Writer
	s struct {
		Status struct {
			Code    string        `json:"code"`
			Number  uint32        `json:"number"`
			Message string        `json:"message"`
			Details []interface{} `json:"details,omitempty"`
		} `json:"status,omitempty"`
		Header   *metadata.MD             `json:"header,omitempty"`
		Messages []map[string]interface{} `json:"messages,omitempty"`
		Trailer  *metadata.MD             `json:"trailer,omitempty"`
	}
	p           present.Presenter
	pbMarshaler *jsonpb.Marshaler
}

func NewResponseFormatter(w io.Writer, emitDefaults bool) format.ResponseFormatterInterface {
	_ = "STUB: not implemented"
	return *new(format.ResponseFormatterInterface)
}

func (p *responseFormatter) FormatHeader(header metadata.MD) { _ = "STUB: not implemented"; return }

func (p *responseFormatter) FormatMessage(v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *responseFormatter) FormatTrailer(trailer metadata.MD) { _ = "STUB: not implemented"; return }

func (p *responseFormatter) FormatStatus(s *status.Status) error {
	_ = "STUB: not implemented"
	return nil
}

// Convert to Any to insert @type field.

func (p *responseFormatter) Done() error { _ = "STUB: not implemented"; return nil }

func (p *responseFormatter) convertProtoMessageToMap(m proto.Message) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *responseFormatter) convertProtoMessageAsAnyToMap(m proto.Message) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
