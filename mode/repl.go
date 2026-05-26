package mode

import (
	"github.com/ktr0731/evans/cache"
	"github.com/ktr0731/evans/config"
	"github.com/ktr0731/evans/cui"
)

func RunAsREPLMode(cfg *config.Config, ui cui.UI, cache *cache.Cache) error {
	_ = "STUB: not implemented"
	return nil
}

func tidyUpHistory(h []string, maxHistorySize int) []string { _ = "STUB: not implemented"; return nil }
