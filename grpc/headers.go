package grpc

// Headers represents gRPC headers. A key corresponds to one or more values.
type Headers map[string][]string

// Add appends a value v to a key k. k must be consisted of other than '-', '_' and '.'.
func (h Headers) Add(k, v string) error {
	_ = "STUB: not implemented"
	// If k is already in h, k is valid key name.
	return nil
}

// Remove deletes values corresponds to a key k.
func (h Headers) Remove(k string) {
	_ = "STUB: not implemented"

	// distinct removes duplicated elements.
	return
}

func distinct(s []string) []string { _ = "STUB: not implemented"; return nil }
