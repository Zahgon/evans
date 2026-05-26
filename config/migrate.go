package config

import (
	"github.com/spf13/viper"
)

var migrationScripts = map[string]func(string, *viper.Viper) string{
	"0.6.10": migrate0610To0611,
}

// migrate migrates an old config schema to the latest one if a new config removes some items that are in an old config.
// Doesn't need to do something if some new items are added; only in cases of deleted or modified old items.
//
// migrate accepts an old version and its config. Migration flow is as follows:
//
//  1. check whether a migration script which migrates the current version to another version. If it is found,
//     apply it to the current config v. Each script is managed by a variable migrationScripts.
//
//  2. The migration script returns a version which is the same as an updated version.
//
//  3. migrate instructs 1 and 2 again with the returned version. If there isn't a new migration script,
//     migrate finishes migration processing.
//
// Note that the latest migration script is not necessarily the same as the Evans's version.
// However, each version of migrationScripts is corresponds to the Evans's one.
func migrate(old string, v *viper.Viper) { _ = "STUB: not implemented"; return }

// No any changes

// migrate0610To0611 migrates a v0.6.10 or older config to v0.6.11 config.
func migrate0610To0611(old string, v *viper.Viper) string { _ = "STUB: not implemented"; return "" }

// request.header in v0.6.10 is formatted like:
//
//   [[request.header]]
//     key = "grpc-client"
//     val = "evans"
//
// These are parsed as a map like:
//
//   []interface {}{
//     {
//       "key": "grpc-client",
//       "val": "evans",
//     },
//   }
//

// v0.6.11 modifies the above structure to a map.

// v0.6.11 renamed Input.PromptFormat to REPL.InputPromptFormat.

// v0.6.11 removed Input field.
