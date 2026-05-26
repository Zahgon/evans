package proto

import (
	"github.com/bufbuild/protocompile/linker"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//go:generate moq -out mock.go . DescriptorSource
type DescriptorSource interface {
	ListServices() ([]string, error)
	FindSymbol(name string) (protoreflect.Descriptor, error)
}

type reflection struct {
	client interface {
		ListServices() ([]string, error)
		FindSymbol(name string) (protoreflect.Descriptor, error)
	}
}

func NewDescriptorSourceFromReflection(c interface {
	ListServices() ([]string, error)
	FindSymbol(name string) (protoreflect.Descriptor, error)
}) DescriptorSource {
	_ = "STUB: not implemented"
	return *new(DescriptorSource)
}

func (r *reflection) ListServices() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *reflection) FindSymbol(name string) (protoreflect.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Descriptor), nil
}

type files struct {
	fds linker.Files
}

func NewDescriptorSourceFromFiles(importPaths []string, fnames []string) (DescriptorSource, error) {
	_ = "STUB: not implemented"
	return *new(DescriptorSource), nil
}

var errSymbolNotFound = errors.New("proto: symbol not found")

func (f *files) ListServices() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *files) FindSymbol(name string) (protoreflect.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Descriptor), nil
}
