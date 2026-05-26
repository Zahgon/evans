// Package cache provides a cache mechanism for the app.
package cache

import (
	"io"
	"os"

	"github.com/ktr0731/evans/meta"
	"github.com/ktr0731/go-updater"
	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
)

const defaultFileName = "cache.toml"

type MeansType updater.MeansType

const MeansTypeUndefined MeansType = ""

type UpdateInfo struct {
	LatestVersion string    `default:"" toml:"latestVersion"`
	InstalledBy   MeansType `default:"" toml:"installedBy"`
}

func (i UpdateInfo) UpdateAvailable() bool { _ = "STUB: not implemented"; return false }

// Cache represents cached items.
type Cache struct {
	Version        string     `toml:"version"`
	UpdateInfo     UpdateInfo `toml:"updateInfo"`
	CommandHistory []string   `default:"" toml:"commandHistory"`

	// SaveFunc is for testing. It will be ignored if it is nil.
	SaveFunc func() error `toml:"-"`
}

// Save writes the receiver to the cache file. It returns an *os.PathError if it can't create a new cache file.
// Also it returns an error if it failed to encode *Cache with TOML format.
func (c *Cache) Save() error { _ = "STUB: not implemented"; return nil }

var decodeTOML = func(r io.Reader, i interface{}) error {
	return toml.NewDecoder(r).Decode(i)
}

// Get returns loaded cache contents.
var Get = func() (*Cache, error) { // Use variable for mocking from means_dev.go.
	p := resolvePath()

	if _, err := os.Stat(p); os.IsNotExist(err) {
		if err := initCacheFile(p); err != nil {
			return nil, errors.Wrap(err, "failed to create a new cache file")
		}
	} else if err != nil {
		return nil, errors.Wrapf(err, "failed to check the existence of cache file '%s'", p)
	}

	f, err := os.Open(p)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open the cache file")
	}
	defer f.Close()

	var c Cache
	if err := decodeTOML(f, &c); err != nil {
		return nil, errors.Wrap(err, "failed to decode loaded cache content")
	}

	// If c.Version is empty or not equal to the latest version, it is regarded as an old version.
	// In such case, we clear the loaded cache.
	if c.Version == "" || c.Version != meta.Version.String() {
		if err := initCacheFile(p); err != nil {
			return nil, errors.Wrap(err, "failed to clear the cache file")
		}
		if _, err := f.Seek(0, 0); err != nil {
			return nil, errors.Wrap(err, "failed to move to the first")
		}
		if err := decodeTOML(f, &c); err != nil {
			return nil, errors.Wrap(err, "failed to decode cache content")
		}
	}

	return &c, nil
}

func resolvePath() string { _ = "STUB: not implemented"; return "" }

// initCacheFile creates or overwrites a new cache file with default values.
// If directories of the file are not found, initCacheFile also creates it.
func initCacheFile(p string) error { _ = "STUB: not implemented"; return nil }
