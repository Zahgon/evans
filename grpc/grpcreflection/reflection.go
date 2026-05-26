// Package grpcreflection provides gRPC reflection client.
// Currently, gRPC reflection depends on Protocol Buffers, so we split this package from grpc package.
package grpcreflection

import (
	"context"

	gr "github.com/jhump/protoreflect/grpcreflect"
	"github.com/ktr0731/grpc-web-go-client/grpcweb"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// ServiceName represents the gRPC reflection service name.
const ServiceName = "grpc.reflection.v1alpha.ServerReflection"

var ErrTLSHandshakeFailed = errors.New("TLS handshake failed")

// Client defines gRPC reflection client.
type Client interface {
	// ListServices lists registered service names.
	// ListServices returns these errors:
	//   - ErrTLSHandshakeFailed: TLS misconfig.
	ListServices() ([]string, error)
	// FindSymbol returns the symbol associated with the given name.
	FindSymbol(name string) (protoreflect.Descriptor, error)
	// Reset clears internal states of Client.
	Reset()
}

type client struct {
	resolver *protoregistry.Files
	client   *gr.Client
}

func getCtx(headers map[string][]string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// NewClient returns an instance of gRPC reflection client for gRPC protocol.
func NewClient(conn grpc.ClientConnInterface, headers map[string][]string) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

// NewWebClient returns an instance of gRPC reflection client for gRPC-Web protocol.
func NewWebClient(conn *grpcweb.ClientConn, headers map[string][]string) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

func (c *client) ListServices() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Check whether the error message contains TLS related error.
// If the server didn't enable TLS, the error message contains the first string.
// If Evans didn't enable TLS against to the TLS enabled server, the error message contains
// the second string.

func (c *client) FindSymbol(name string) (protoreflect.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Descriptor), nil
}

// TODO: consider dependencies

func (c *client) Reset() { _ = "STUB: not implemented"; return }
