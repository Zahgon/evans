package usecase

// UseService modifies svcName as the currently selected service.
// UseService may return these errors:
//
//   - ErrPackageUnselected: REPL never call UsePackage.
//   - ErrUnknownServiceName: svcName is not in loaded services.
func UseService(svcName string) error { _ = "STUB: not implemented"; return nil }

func (m *dependencyManager) UseService(svcName string) error { _ = "STUB: not implemented"; return nil }

// Keep backward-compatibility.
// TODO: Delete package related code after releasing v1.0.0.

// In the case of empty package.
