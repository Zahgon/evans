package usecase

import (
	"github.com/ktr0731/evans/grpc"
)

// ListRPCs lists all RPC belong to the selected service.
// If svcName is empty, the currently selected service will be used.
// In this case, ListRPCs doesn't modify the currently selected service.
func ListRPCs(svcName string) ([]*grpc.RPC, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *dependencyManager) ListRPCs(svcName string) ([]*grpc.RPC, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *dependencyManager) listRPCs(fqsn string) ([]*grpc.RPC, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: handle "ok".
