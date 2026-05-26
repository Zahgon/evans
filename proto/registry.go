package proto

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

type anyResolver struct {
	protoregistry.ExtensionTypeResolver
	descSource DescriptorSource
}

func NewAnyResolver(descSource DescriptorSource) interface {
	protoregistry.ExtensionTypeResolver
	protoregistry.MessageTypeResolver
} {
	_ = "STUB: not implemented"
	return nil
}

func (r *anyResolver) FindMessageByName(m protoreflect.FullName) (protoreflect.MessageType, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.MessageType), nil
}

// Fallback to protoregistry.GlobalTypes.

// TODO: handle "ok".

func (r *anyResolver) FindMessageByURL(url string) (protoreflect.MessageType, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.MessageType), nil
}

// Fallback to protoregistry.GlobalTypes.

// TODO: handle "ok".
