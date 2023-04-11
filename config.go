package setting

import (
	"github.com/spf13/viper"
)

// Config interface associated with reading/saving configuration files.
type Config interface {
	// Load config load
	Load() error

	// IsSet is set value
	IsSet(key string) bool

	// ReadSection read val by key,val must be a pointer
	ReadSection(key string, val interface{}) error

	// Store save config to file
	Store(path string) error
}

// New create a config interface.
func New(opts ...Option) Config {
	c := &viperConfig{
		vp:       viper.New(),
		sections: make(map[string]interface{}, 20),
	}

	conf := &Options{
		configFile: "./app.yaml",
	}

	for _, o := range opts {
		o(conf)
	}

	c.configFile = conf.configFile
	c.watchFile = conf.watchFile

	return c
}
