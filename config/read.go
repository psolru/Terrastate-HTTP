package config

import (
	"github.com/psolru/terrastate-http/env"
	"os"
)

// getConfigFile loads the actual config file from disk
func getConfigFile() (*os.File, error) {
	return os.Open(env.WorkDir() + "/configs/config.json")
}
