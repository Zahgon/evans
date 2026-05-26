package grpc

import (
	"context"

	"github.com/ktr0731/evans/grpc/grpcreflection"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var ErrMutualAuthParamsAreNotEnough = errors.New("cert and certkey are required to authenticate mutually")

// RPC represents a RPC which belongs to a gRPC service.
type RPC struct {
	Name               string
	FullyQualifiedName string
	RequestType        *Type
	ResponseType       *Type
	IsServerStreaming  bool
	IsClientStreaming  bool
}

// Type is a type for representing requests/responses.
type Type struct {
	// Name is the name of Type.
	Name string

	// FullyQualifiedName is the name that contains the package name this Type belongs.
	FullyQualifiedName string

	// New instantiates a new instance of Type.  It is used for decode requests and responses.
	New func() interface{}
}

// Client represents the gRPC client.
type Client interface {
	// Invoke invokes a request req to the gRPC server. Then, Invoke decodes the response to res.
	Invoke(ctx context.Context, fqrn string, req, res interface{}) (header, trailer metadata.MD, _ error)

	// NewClientStream creates a new client stream.
	NewClientStream(ctx context.Context, streamDesc *grpc.StreamDesc, fqrn string) (ClientStream, error)

	// NewServerStream creates a new server stream.
	NewServerStream(ctx context.Context, streamDesc *grpc.StreamDesc, fqrn string) (ServerStream, error)

	// NewBidiStream creates a new bidirectional stream.
	NewBidiStream(ctx context.Context, streamDesc *grpc.StreamDesc, fqrn string) (BidiStream, error)

	// Close closes all connections the client has.
	Close(ctx context.Context) error

	// Header returns all request headers (metadata) Client has.
	Header() Headers

	grpcreflection.Client
}

type ClientStream interface {
	// Header returns the response header.
	Header() (metadata.MD, error)
	// Trailer returns the response trailer.
	Trailer() metadata.MD
	Send(req interface{}) error
	CloseAndReceive(res interface{}) error
}

type ServerStream interface {
	// Header returns the response header.
	Header() (metadata.MD, error)
	// Trailer returns the response trailer.
	Trailer() metadata.MD
	Send(req interface{}) error
	Receive(res interface{}) error
}

type BidiStream interface {
	// Header returns the response header.
	Header() (metadata.MD, error)
	// Trailer returns the response trailer.
	Trailer() metadata.MD
	Send(req interface{}) error
	Receive(res interface{}) error
	CloseSend() error
}

type client struct {
	conn    *grpc.ClientConn
	headers Headers

	grpcreflection.Client
}

// NewClient creates a new gRPC client. It dials to the server specified by addr.
// addr format is the same as the first argument of grpc.Dial.
// If serverName is not empty, it overrides the gRPC server name used to
// verify the hostname on the returned certificates.
// If useReflection is true, the gRPC client enables gRPC reflection.
// If useTLS is true, the gRPC client establishes a secure connection with the server.
//
// The set of cert and certKey enables mutual authentication if useTLS is enabled.
// If one of it is not found, NewClient returns ErrMutualAuthParamsAreNotEnough.
// If useTLS is false, cacert, cert and certKey are ignored.
func NewClient(addr, serverName string, useReflection, useTLS bool, cacert, cert, certKey string, headers map[string][]string) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

// Enable TLS authentication

// Enable mutual authentication

func (c *client) Invoke(ctx context.Context, fqrn string, req, res interface{}) (header, trailer metadata.MD, _ error) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), *new(metadata.MD), nil
}

func (c *client) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *client) Header() Headers { _ = "STUB: not implemented"; return *new(Headers) }

type clientStream struct {
	cs grpc.ClientStream
}

func (s *clientStream) Header() (metadata.MD, error) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), nil
}

func (s *clientStream) Trailer() metadata.MD { _ = "STUB: not implemented"; return *new(metadata.MD) }

func (s *clientStream) Send(req interface{}) error { _ = "STUB: not implemented"; return nil }

func (s *clientStream) CloseAndReceive(res interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) NewClientStream(ctx context.Context, streamDesc *grpc.StreamDesc, fqrn string) (ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(ClientStream), nil
}

type serverStream struct {
	*clientStream
}

func (s *serverStream) Receive(res interface{}) error { _ = "STUB: not implemented"; return nil }

func (c *client) NewServerStream(ctx context.Context, streamDesc *grpc.StreamDesc, fqrn string) (ServerStream, error) {
	_ = "STUB: not implemented"
	return *new(ServerStream), nil
}

type bidiStream struct {
	s *serverStream
}

func (s *bidiStream) Header() (metadata.MD, error) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), nil
}

func (s *bidiStream) Trailer() metadata.MD { _ = "STUB: not implemented"; return *new(metadata.MD) }

func (s *bidiStream) Send(req interface{}) error { _ = "STUB: not implemented"; return nil }

func (s *bidiStream) Receive(res interface{}) error { _ = "STUB: not implemented"; return nil }

func (s *bidiStream) CloseSend() error { _ = "STUB: not implemented"; return nil }

func (c *client) NewBidiStream(ctx context.Context, streamDesc *grpc.StreamDesc, fqrn string) (BidiStream, error) {
	_ = "STUB: not implemented"
	return *new(BidiStream), nil
}

// fqrnToEndpoint converts FullQualifiedRPCName to endpoint
//
// e.g.
//
//	pkg_name.svc_name.rpc_name -> /pkg_name.svc_name/rpc_name
func fqrnToEndpoint(fqrn string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// FQRN should contain at least service and rpc name.

func wakeUpClientConn(conn *grpc.ClientConn) { _ = "STUB: not implemented"; return }

func loggingRequest(req interface{}) { _ = "STUB: not implemented"; return }
