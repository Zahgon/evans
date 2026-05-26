package usecase

// UsePackage modifies pkgName as the currently selected package.
// UsePackage may return these errors:
//
//   - ErrUnknownPackageName: pkgName is not in loaded packages.
func UsePackage(pkgName string) error { _ = "STUB: not implemented"; return nil }

func (m *dependencyManager) UsePackage(pkgName string) error { _ = "STUB: not implemented"; return nil }
