package app

// flags defines available command line flags.
//
//nolint:maligned
type flags struct {
	mode struct {
		repl bool
		cli  bool
	}

	cli struct {
		call string
		file string
	}

	repl struct {
		silent bool
	}

	common struct {
		pkg        string
		service    string
		path       []string
		proto      []string
		host       string
		port       string
		header     map[string][]string
		web        bool
		reflection bool
		tls        bool
		cacert     string
		cert       string
		certKey    string
		serverName string
	}

	meta struct {
		edit       bool
		editGlobal bool
		verbose    bool
		version    bool
		help       bool
	}
}

// validate defines invalid conditions and validates whether f has invalid conditions.
func (f *flags) validate() error { _ = "STUB: not implemented"; return nil }

// -- stringToString Value.
type stringToStringSliceValue struct {
	value   *map[string][]string
	changed bool
}

func newStringToStringValue(val map[string][]string, p *map[string][]string) *stringToStringSliceValue {
	_ = "STUB: not implemented"
	return nil
}

// Format: a=1,b=2.
func (s *stringToStringSliceValue) Set(val string) error { _ = "STUB: not implemented"; return nil }

func (s *stringToStringSliceValue) Type() string { _ = "STUB: not implemented"; return "" }

func (s *stringToStringSliceValue) String() string { _ = "STUB: not implemented"; return "" }
