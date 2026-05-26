package prompt

type opt struct {
	commandHistory []string
}

type Option func(*opt)

func WithCommandHistory(h []string) Option { _ = "STUB: not implemented"; return *new(Option) }
