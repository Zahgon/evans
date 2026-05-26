package proto

import (
	"github.com/ktr0731/evans/fill"
	"github.com/ktr0731/evans/prompt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

// InteractiveFiller is an implementation of fill.InteractiveFiller.
// It let you input request fields interactively.
type InteractiveFiller struct {
	prompt       prompt.Prompt
	prefixFormat string
}

// NewInteractiveFiller instantiates a new filler that fills each field interactively.
func NewInteractiveFiller(prompt prompt.Prompt, prefixFormat string) *InteractiveFiller {
	_ = "STUB: not implemented"
	return nil
}

// Fill receives v that is an instance of *dynamic.Message.
// Fill let you input each field interactively by using a prompt. v will be set field values inputted by a prompt.
//
// Note that Fill resets the previous state when it is called again.
func (f *InteractiveFiller) Fill(v *dynamicpb.Message, opts fill.InteractiveFillerOpts) error {
	_ = "STUB: not implemented"
	return nil
}

type resolver struct {
	prompt       prompt.Prompt
	prefixFormat string
	color        prompt.Color

	msg *dynamicpb.Message

	m         protoreflect.MessageDescriptor
	ancestors []string
	// repeated represents that the message is repeated field or not.
	// If the message is not a field or not a repeated field, it is false.
	repeated bool

	opts fill.InteractiveFillerOpts
}

func newResolver(
	prompt prompt.Prompt,
	prefixFormat string,
	color prompt.Color,
	msg *dynamicpb.Message,
	ancestors []string,
	repeated bool,
	opts fill.InteractiveFillerOpts,
) *resolver {
	_ = "STUB: not implemented"
	return nil
}

func (r *resolver) resolve() (*dynamicpb.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for _, f := range r.m.Fields(). {

// Skip if one of choices is already selected.

func (r *resolver) resolveOneof(o protoreflect.OneofDescriptor) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *resolver) resolveField(f protoreflect.FieldDescriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// For bytes, if neither BytesAsBase64 nor BytesAsQuotedLiterals is explicitly set,
// try to decode as base64 first, and if that fails, fall back trying to parse
// as quoted literals (logging a warning).
//
// This is to preserve backwards compatibility, as we used to be accept quoted
// literals, but want to switch to base64.
//
// If either BytesAsBase64 or BytesAsQuotedLiterals is set, only parse it in that format,
// and if BytesFromFile is set, read it from file.
//
// Use strconv.Unquote to interpret byte literals and Unicode literals.
// For example, a user inputs `\x6f\x67\x69\x73\x6f`,
// His expects "ogiso" in string, but backslashes in the input are not interpreted as an escape sequence.
// So, we need to call strconv.Unquote to interpret backslashes as an escape sequence.

// try to decode as base64

// failed, try to parse as quoted literal

// failed to parse as this too, assume user intended to input base64, propagate
// that error

// log a warning and return the decoded literal string

// succeeded decoding as base64, return

// TODO: or cardinality

// TODO: is it okay?

// Return nil to keep inputted values.

// io.EOF signals the end of inputting repeated field.
// Return nil to keep inputted values.

func (r *resolver) resolveEnum(prefix string, e protoreflect.EnumDescriptor) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// for _, v := range e.GetValues() {

func (r *resolver) input(prefix string, f protoreflect.FieldDescriptor, converter func(string) (protoreflect.Value, error)) (protoreflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Value), nil
}

func (r *resolver) selectChoices(msg string, choices []string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Skip inputting and use default value.

func (r *resolver) addRepeatedField(f protoreflect.FieldDescriptor) bool {
	_ = "STUB: not implemented"
	return false
}

// f is repeated empty message field. It will cause infinite-loop if r.opts.AddRepeatedManually is false.
// For user's experience, always display prompt in this case.

func (r *resolver) skipMessage(f protoreflect.FieldDescriptor) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *resolver) makePrefix(field protoreflect.FieldDescriptor) string {
	_ = "STUB: not implemented"
	return ""
}

var protoDefaults = map[protoreflect.Kind]interface{}{
	protoreflect.DoubleKind:   float64(0),
	protoreflect.FloatKind:    float32(0),
	protoreflect.Int64Kind:    int64(0),
	protoreflect.Uint64Kind:   uint64(0),
	protoreflect.Int32Kind:    int32(0),
	protoreflect.Uint32Kind:   uint32(0),
	protoreflect.Fixed64Kind:  uint64(0),
	protoreflect.Fixed32Kind:  uint32(0),
	protoreflect.BoolKind:     false,
	protoreflect.StringKind:   "",
	protoreflect.BytesKind:    []byte{},
	protoreflect.Sfixed64Kind: int64(0),
	protoreflect.Sfixed32Kind: int32(0),
	protoreflect.Sint64Kind:   int64(0),
	protoreflect.Sint32Kind:   int32(0),
}

// convertValue converts a string input pv to protoreflect.Value.
func defaultValueFromKind(kind protoreflect.Kind) protoreflect.Value {
	_ = "STUB: not implemented"
	return *new(protoreflect.Value)
}
