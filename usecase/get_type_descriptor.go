package usecase

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

// GetTypeDescriptor gets the descriptor of a type which belongs to the currently selected package.
func GetTypeDescriptor(typeName string) (protoreflect.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Descriptor), nil
}

func (m *dependencyManager) GetTypeDescriptor(typeName string) (protoreflect.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Descriptor), nil
}
