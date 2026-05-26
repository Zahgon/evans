package usecase

import (
	"context"
	"io"

	"github.com/ktr0731/evans/fill"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

// ErrorCode represents an application error code.
type ErrorCode int

// String implements fmt.Stringer.
func (e ErrorCode) String() string { _ = "STUB: not implemented"; return "" }

type gRPCError struct {
	*status.Status
}

func (e *gRPCError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *gRPCError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *gRPCError) Code() ErrorCode { _ = "STUB: not implemented"; return *new(ErrorCode) }

// CallRPC constructs a request with input source such that prompt inputting, stdin or a file. After that, it sends
// the request to the gRPC server and decodes the response body to res.
// Note that req and res must be JSON-decodable structs. The output is written to w.
func CallRPC(ctx context.Context, w io.Writer, rpcName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *dependencyManager) CallRPC(ctx context.Context, w io.Writer, rpcName string, rerunPrevious bool, filler fill.Filler) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: handle "ok".

// Client streaming RPCs are RPC that a client sends several times and server responds once.
// Client streaming RPCs are processed by the following instruction.
//
//   1. Create a new client stream.
//   2. Create a new request and fill input to it.
//   3. Send the request to the server.
//   4. Repeat 1-3 until the filler returns io.EOF.
//   5. Send the close message and receive the response.
//   6. Format the response and output it.
//

// gRPC error. Treat as a normal response.

// Server streaming RPCs are RPC that a client sends once and server responds several times.
// Server streaming RPCs are processed by the following instruction.
//
//   1. Create a new request and fill input to it.
//   2. Send the request to the server.
//   3. Call Receive to receive server responses.
//   4. Format a received response and output it.
//   5. If io.EOF received, finish the RPC connection.
//

// Trailer is now available.

// If both of rpc.IsStreamingClient() and rpc.IsStreamingServer() are false, it means its RPC is an unary RPC.
// Unary RPCs are processed by the following instruction.
//
//   1. Create a new request and fill input to it.
//   2. Create a new response.
//   3. Invoke the RPC with the request and decode response to 2's instance.
//   4. Format the response.
//

// Gets a request with the body containing the payload of its previous method
// Only RPCs that are repeatable by definition will have their previous requests returned.
func (m *dependencyManager) getPreviousRPCRequest(method protoreflect.MethodDescriptor, req proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Custom resolver.

// Updates the last call state for the given method. This is done by serializing
// the request payload and store it into the state buffer indexed by the rpcName
// The method is repeatable only if it is not a client streaming method.
func (m *dependencyManager) updateMethodCallState(rpcName string, req *dynamicpb.Message) error {
	_ = "STUB: not implemented"
	return nil
}

type interactiveFiller struct {
	fillFunc func(v *dynamicpb.Message) error
}

func (f *interactiveFiller) Fill(v *dynamicpb.Message) error { _ = "STUB: not implemented"; return nil }

func CallRPCInteractively(ctx context.Context, w io.Writer, rpcName string, digManually, bytesAsBase64, bytesAsQuotedLiterals, bytesFromFile, rerunPrevious, addRepeatedManually bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *dependencyManager) CallRPCInteractively(ctx context.Context, w io.Writer, rpcName string, digManually, bytesAsBase64, bytesAsQuotedLiterals, bytesFromFile, rerunPrevious, addRepeatedManually bool) error {
	_ = "STUB: not implemented"
	return nil
}

func handleGRPCResponseError(err error) (*status.Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
