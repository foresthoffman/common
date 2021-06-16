// Package common provides helpful DRY logic common to most Go services.
package common

import "os"

// GetEnv returns the provided environment variable, if set. The value defaults
// to the fallback, if the variable is empty.
func GetEnv(name string, fallback string) string {
	v := os.Getenv(name)
	if v == "" {
		return fallback
	}
	return v
}
