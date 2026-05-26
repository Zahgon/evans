package usecase

// ParseFullyQualifiedMethodName parses the passed fully-qualified method as fully-qualified service name and method name.
// ParseFullyQualifiedMethodName may return these errors:
//
//   - An error described in idl.Spec.RPC method returns.
//   - An error if fqmn is not a valid fully-qualified method name form.
func ParseFullyQualifiedMethodName(fqmn string) (fqsn, method string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func (m *dependencyManager) ParseFullyQualifiedMethodName(fqmn string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
