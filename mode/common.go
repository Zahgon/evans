package mode

import (
	"github.com/ktr0731/evans/config"
	"github.com/ktr0731/evans/grpc"
	"github.com/ktr0731/evans/grpc/grpcreflection"
	"github.com/ktr0731/evans/proto"
)

func newGRPCClient(cfg *config.Config) (grpc.Client, error) {
	_ = "STUB: not implemented"
	return *new(grpc.Client), nil
}

//TODO: remove second arg

func gRPCReflectionPackageFilteredPackages(pkgNames []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func setDefault(cfg *config.Config) error {
	_ = "STUB: not implemented"
	// If the spec has only one package, mark it as the default package.
	return nil
}

// If the spec has only one service, mark it as the default service.

// Ignore server reflection name because it's provided imply when reflection is enabled.

func newDescSource(cfg *config.Config, grpcClient grpcreflection.Client) (descSource proto.DescriptorSource, err error) {
	_ = "STUB: not implemented"
	return *new(proto.DescriptorSource), nil
}

func dropString(slice []string, s string) []string { _ = "STUB: not implemented"; return nil }
