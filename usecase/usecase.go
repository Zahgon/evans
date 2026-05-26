// Package usecase provides all use-case logics used from each mode.
// Clients of usecase package must call Inject before all function calls in usecase package.
package usecase

import (
	"github.com/ktr0731/evans/fill"
	"github.com/ktr0731/evans/format"
	"github.com/ktr0731/evans/grpc"
	"github.com/ktr0731/evans/present"
	"github.com/ktr0731/evans/proto"
	"github.com/pkg/errors"
)

var (
	ErrPackageUnselected = errors.New("package unselected")
	ErrServiceUnselected = errors.New("service unselected")

	ErrUnknownPackageName = errors.New("unknown package name")
	ErrUnknownServiceName = errors.New("unknown service name")
	ErrUnknownRPCName     = errors.New("unknown RPC name")
	ErrUnknownSymbol      = errors.New("unknown symbol")
)

var (
	defaultState = state{}
	dm           = &dependencyManager{}
)

type dependencyManager struct {
	descSource        proto.DescriptorSource
	filler            fill.Filler
	interactiveFiller fill.InteractiveFiller
	gRPCClient        grpc.Client
	responseFormatter *format.ResponseFormatter
	resourcePresenter present.Presenter
	state             state
}

type rpcIdentifier string

// state has the domain state modified by each usecase logic. The default value is used as the initial value.
type state struct {
	selectedPackage string // TODO: remove in v1.0.0.
	selectedService string
	rpcCallState    map[rpcIdentifier]callState
}

type callState struct {
	requestPayload []byte
}

type Dependencies struct {
	DescSource        proto.DescriptorSource
	Filler            fill.Filler
	InteractiveFiller fill.InteractiveFiller
	GRPCClient        grpc.Client
	ResponseFormatter *format.ResponseFormatter
	ResourcePresenter present.Presenter
}

// Inject corresponds an implementation to an interface type. Inject clears the previous states if it exists.
func Inject(deps Dependencies) { _ = "STUB: not implemented"; return }

func (m *dependencyManager) Inject(d Dependencies) { _ = "STUB: not implemented"; return }

// InjectPartially is almost same as the Inject, but injects only non-nil dependencies.
func InjectPartially(deps Dependencies) { _ = "STUB: not implemented"; return }

func (m *dependencyManager) InjectPartially(d Dependencies) { _ = "STUB: not implemented"; return }

// Clear clears all dependencies and states. Usually, it is used for unit testing.
func Clear() { _ = "STUB: not implemented"; return }
