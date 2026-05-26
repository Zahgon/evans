// Package config provides config structures, and a mechanism that merges sources such that the global config file,
// a project local config file and command line flags.
package config

import (
	"os"
	"os/exec"

	"github.com/ktr0731/go-multierror"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	localConfigName  = ".evans.toml"
	globalConfigName = "config.toml"
)

type Server struct {
	Host       string `toml:"host"`
	Port       string `toml:"port"`
	Reflection bool   `toml:"reflection"`
	TLS        bool   `toml:"tls"`
	Name       string `toml:"name"`
}

type Header map[string][]string

type Request struct {
	Header      Header `toml:"header"`
	Web         bool   `toml:"web"`
	CACertFile  string `toml:"caCertFile"`
	CertFile    string `toml:"certFile"`
	CertKeyFile string `toml:"certKeyFile"`
}

type REPL struct {
	PromptFormat      string `toml:"promptFormat"`
	InputPromptFormat string `toml:"inputPromptFormat"`

	ColoredOutput bool `toml:"coloredOutput"`

	Silent         bool   `toml:"silent"`
	SplashTextPath string `toml:"splashTextPath"`

	// TODO: Split history files between projects.
	HistorySize int `toml:"historySize"`
}

type Meta struct {
	ConfigVersion string `toml:"configVersion"`
	AutoUpdate    bool   `toml:"autoUpdate"`
	UpdateLevel   string `toml:"updateLevel"`
}

// Each TOML key must be equal the field name in the lower-case. It is a limitation of spf13/viper.
type Config struct {
	Default *Default `toml:"default"`
	Meta    *Meta    `toml:"meta"`
	REPL    *REPL    `toml:"repl"`
	Server  *Server  `toml:"server"`
	Log     *Log     `toml:"log"`
	Request *Request `toml:"request"`
}

// ValidationError contains errors that describes invalid config conditions.
type ValidationError struct {
	Err *multierror.Error
}

// Error returns ValidationError's error text.
func (e *ValidationError) Error() string { _ = "STUB: not implemented"; return "" }

// Validate defines invalid conditions and validates whether c has invalid condtions.
// For example, in the case of CLI mode, c must have package, service and call values.
// Validate returns ValidationError if some conditions are invalid.
func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// TODO: support it.

type Default struct {
	ProtoPath []string `toml:"protoPath"`
	ProtoFile []string `toml:"protoFile"`
	Package   string   `toml:"package"`
	Service   string   `toml:"service"`
}

type Log struct {
	Prefix string `toml:"prefix"`
}

// Get returns the config which loaded from the global and local config files,
// and command line flags passed as an argument. Note that fs must have been parsed.
//
// The order of priority is flags > local > global.
func Get(fs *pflag.FlagSet) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

func newDefaultViper() *viper.Viper { _ = "STUB: not implemented"; return nil }

// We set the default version to v0.6.10 because the structure of Config is changed at v0.6.11.

// bindFlags binds parsed flag values to vp. Note that fs must be parsed.
func bindFlags(vp *viper.Viper, fs *pflag.FlagSet) {
	_ = "STUB: not implemented"
	// kv defines the mapping from a viper config name to a flag name.
	return
}

// There is pflag.StringToString which converts 'key=val' to a map structure.
// However, currently, we don't use BindPFlag because it has some bugs.

// We want to append flag values to the config.
// So, we don't use BindPFlag.

// stringToStringSliceToMap converts (app.stringToStringSliceValue).String() to a map.
// If some errors occur, stringToStringSliceToMap returns an empty map.
func stringToStringSliceToMap(val string) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

// stringToStringToMap converts (pflag.stringToStringValue).String() to a map.
// If some errors occur, stringToStringToMap returns an empty map.
func stringToStringToMap(val string) map[string]string { _ = "STUB: not implemented"; return nil }

func stringSliceToSlice(val string) []string { _ = "STUB: not implemented"; return nil }

// writeLatestDefaultConfig writes the latest default config to path.
// Note that writeLatestDefaultConfig initializes viper again.
// So, all flags you bind by BindPFlag, global and local config will be clear.
func writeLatestDefaultConfig(path string) (*Config, error) {
	_ = "STUB: not implemented"
	return nil,

		// Set configVersion to the latest version.
		nil
}

func initConfig(fs *pflag.FlagSet) (cfg *Config, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Global config paths

// Migrate old versions to the latest.

// Update the global config with the migrated config.

func setupConfig(c *Config) {
	_ = "STUB: not implemented"
	// To show protofile and protopath field in a config file, set slice which has empty string
	// if these are nil. (please see default values.)
	// Conversely, trim the empty string element when config loading.
	return
}

// Edit opens the project local config file with an editor.
// If the local config file is missing, Edit creates a new local config file.
// $EDITOR is used as an editor if it is configured. Else, Vim is used.
func Edit() error { _ = "STUB: not implemented"; return nil }

// EditGlobal is the same as Edit, but edit the global config.
func EditGlobal() error { _ = "STUB: not implemented"; return nil }

var runEditor = func(editor string, cfgPath string) error {
	cmd := exec.Command(editor, cfgPath)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "failed to execute %s", editor)
	}
	return nil
}

func getLocalConfigPath() (string, bool) { _ = "STUB: not implemented"; return "", false }

// getGlobalConfig returns always the global config path.
// If the file is missing, it returns false as the second returned value.
func getGlobalConfigPath() (string, bool) { _ = "STUB: not implemented"; return "", false }

func lookupProjectRootPath() (string, bool) { _ = "STUB: not implemented"; return "", false }

func getEditor() string { _ = "STUB: not implemented"; return "" }
