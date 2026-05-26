package usecase

import (
	"github.com/ktr0731/evans/grpc"
)

func AddHeader(k, v string) { _ = "STUB: not implemented"; return }

func (m *dependencyManager) AddHeader(k, v string) { _ = "STUB: not implemented"; return }

func RemoveHeader(k string) { _ = "STUB: not implemented"; return }

func (m *dependencyManager) RemoveHeader(k string) { _ = "STUB: not implemented"; return }

func ListHeaders() grpc.Headers { _ = "STUB: not implemented"; return *new(grpc.Headers) }

func (m *dependencyManager) ListHeaders() grpc.Headers {
	_ = "STUB: not implemented"
	return *new(grpc.Headers)
}
