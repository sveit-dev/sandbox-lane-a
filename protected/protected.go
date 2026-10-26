// Package protected holds the values the rest of the module depends on.
// Changes here need a human review; agents must not edit this directory.
package protected

import (
	_ "embed"
	"encoding/json"
)

// MaxItems is the largest total quantity a cart may hold.
const MaxItems = 50

//go:embed config.json
var raw []byte

// Config is the parsed content of config.json.
type Config struct {
	Currency        string `json:"currency"`
	DefaultPageSize int    `json:"default_page_size"`
}

// Load parses the embedded config.json.
func Load() (Config, error) {
	var c Config
	err := json.Unmarshal(raw, &c)
	return c, err
}
