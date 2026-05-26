package usecase

// FormatMethods formats all method names.
func FormatMethods() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (m *dependencyManager) FormatMethods() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *dependencyManager) methodsToFormatStructs(fqsn string) (v struct {
	Methods []struct {
		Name               string `json:"name" table:"name"`
		FullyQualifiedName string `json:"fully_qualified_name" name:"target" table:"fully-qualified name"`
		RequestType        string `json:"request_type" table:"request type"`
		ResponseType       string `json:"response_type" table:"response type"`
	} `json:"methods" name:"target"`
}, _ error) {
	_ = "STUB: not implemented"
	return nil, nil
}
