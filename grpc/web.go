package grpc

import (
	"context"

	"github.com/ktr0731/evans/grpc/grpcreflection"
	"github.com/ktr0731/grpc-web-go-client/grpcweb"
	gogrpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type webClient struct {
	conn    *grpcweb.ClientConn
	headers Headers

	grpcreflection.Client
}

func NewWebClient(addr string, useReflection, useTLS bool, cacert, cert, certKey string, headers Headers) Client {
	_ = "STUB: not implemented"
	return *new(Client)
}

func (c *webClient) Invoke(ctx context.Context, fqrn string, req, res interface{}) (header, trailer metadata.MD, _ error) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), *new(metadata.MD), nil
}

type webClientStream struct {
	ctx    context.Context
	stream grpcweb.ClientStream
}

func (s *webClientStream) Header() (metadata.MD, error) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), nil
}

func (s *webClientStream) Trailer() metadata.MD {
	_ = "STUB: not implemented"
	return *new(metadata.MD)
}

func (s *webClientStream) Send(req interface{}) error { _ = "STUB: not implemented"; return nil }

func (s *webClientStream) CloseAndReceive(res interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *webClient) NewClientStream(ctx context.Context, streamDesc *gogrpc.StreamDesc, fqrn string) (ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(ClientStream), nil
}

type webServerStream struct {
	ctx    context.Context
	stream grpcweb.ServerStream
}

func (s *webServerStream) Header() (metadata.MD, error) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), nil
}

func (s *webServerStream) Trailer() metadata.MD {
	_ = "STUB: not implemented"
	return *new(metadata.MD)
}

func (s *webServerStream) Send(req interface{}) (err error) { _ = "STUB: not implemented"; return nil }

func (s *webServerStream) Receive(res interface{}) error { _ = "STUB: not implemented"; return nil }

func (c *webClient) NewServerStream(ctx context.Context, streamDesc *gogrpc.StreamDesc, fqrn string) (ServerStream, error) {
	_ = "STUB: not implemented"
	return *new(ServerStream), nil
}

type webBidiStream struct {
	ctx    context.Context
	stream grpcweb.BidiStream
}

func (s *webBidiStream) Header() (metadata.MD, error) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), nil
}

func (s *webBidiStream) Trailer() metadata.MD { _ = "STUB: not implemented"; return *new(metadata.MD) }

func (s *webBidiStream) Send(req interface{}) error { _ = "STUB: not implemented"; return nil }

func (s *webBidiStream) Receive(res interface{}) error { _ = "STUB: not implemented"; return nil }

func (s *webBidiStream) CloseSend() error { _ = "STUB: not implemented"; return nil }

func (c *webClient) NewBidiStream(ctx context.Context, streamDesc *gogrpc.StreamDesc, fqrn string) (BidiStream, error) {
	_ = "STUB: not implemented"
	return *new(BidiStream), nil
}

func (c *webClient) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *webClient) Header() Headers { _ = "STUB: not implemented"; return *new(Headers) }
