// Package curl provides a curl-like formatter implementation.
package curl

import (
	"io"
	"strings"

	"github.com/golang/protobuf/jsonpb" //nolint:staticcheck
	"github.com/golang/protobuf/proto"  //nolint:staticcheck
	"github.com/ktr0731/evans/format"
	"github.com/ktr0731/evans/present"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails" // For calling RegisterType.
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type responseFormatter struct {
	w io.Writer

	json        present.Presenter
	pbMarshaler *jsonpb.Marshaler

	wroteHeader, wroteMessage, wroteTrailer bool
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

var replacer = strings.NewReplacer("\n", "", ",", ", ")

func (p *responseFormatter) FormatStatus(status *status.Status) error {
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
