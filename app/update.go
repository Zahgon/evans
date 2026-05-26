package app

import (
	"context"
	"io"
	"syscall"

	"github.com/hashicorp/go-version"
	"github.com/ktr0731/evans/cache"
	"github.com/ktr0731/evans/config"
	"github.com/ktr0731/evans/prompt"
	"github.com/ktr0731/go-updater"
)

// checkUpdate checks whether an update exists. Update checking is instructed by following steps:
//  1. If install means is known, use it as an update means.
//     If install means is unknown, checkUpdate selects an available means from candidates.
//  2. Check whether update exists. If it is found, cache the latest version.
func checkUpdate(ctx context.Context, cfg *config.Config, c *cache.Cache) error {
	_ = "STUB: not implemented"
	return nil
}

// if ErrUnavailable, user installed Evans by manually, ignore.

// If specified means builder is not found, skip update.

var syscallExec = syscall.Exec

// processUpdate checks new changes and updates Evans in accordance with user's selection.
// If config.Meta.AutoUpdate enabled, processUpdate is called asynchronously.
// Other than, processUpdate is called synchronously.
func processUpdate(ctx context.Context, cfg *config.Config, w io.Writer, c *cache.Cache, prompt prompt.Prompt) error {
	_ = "STUB: not implemented"
	return nil
}

// If cached version is less than or equal to current version, ignore it.

// Instantiate the means.
// If ErrUnavailable, user installed Evans by manually, ignore.

// If auto update is enabled, do process update without user's confirmation.

// If canceled, ignore and return

// If auto update is disabled, display update info

// Abort updating.

// If canceled, ignore and return

// restart Evans

// update updates Evans to the latest version. If interrupted by a key, update will be canceled.
func update(ctx context.Context, infoWriter io.Writer, updater *updater.Updater, c *cache.Cache) error {
	_ = "STUB: not implemented"
	return nil
}

// update successful

var updateInfoFormat = `
new update available:
  current version: %s
   latest version: %s

`

func printUpdateInfo(w io.Writer, latest string) { _ = "STUB: not implemented"; return }

// newUpdater creates new updater from cached information. updater checks whether UpdateIf is true or false
// to display update information to the user.
func newUpdater(cfg *config.Config, v *version.Version, m updater.Means) *updater.Updater {
	_ = "STUB: not implemented"
	return nil
}
