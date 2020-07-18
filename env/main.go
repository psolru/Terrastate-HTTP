package env

import "os"

// GetEnv returns the value of the given environment variable key
// If it can't find the key it will use the provided fallback
// the fallback must be given
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// MountPoint returns the location where work files are provided to go
func WorkDir() string {
	return GetEnv("WORKDIR", "/mnt")
}
