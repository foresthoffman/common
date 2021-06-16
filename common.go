// Package common provides helpful DRY logic common to most Go services.
package common

import (
	"fmt"
	"os"
	"strings"
)

// GetEnv returns the provided environment variable, if set. The value defaults
// to the fallback, if the variable is empty.
func GetEnv(name string, fallback string) string {
	v := os.Getenv(name)
	if v == "" {
		return fallback
	}
	return v
}

// Dash concatenates the provided strings with a dash delimiter (a standard hyphen: "-").
// Empty segments and segments without non-delimiter characters will be ignored.
func Dash(segments ...string) string {
	var out string
	for _, segment := range segments {
		fullTrimmed := strings.Trim(segment, " -")
		if len(fullTrimmed) == 0 {
			continue
		}
		trimmedOut := strings.Trim(out, "-")
		if len(trimmedOut) > 0 {
			out = fmt.Sprintf("%s-%s", trimmedOut, fullTrimmed)
			continue
		}
		out = fullTrimmed
	}
	return out
}

// Url combines the path segments with backslashes regardless of whether or not
// the segments are prefixed/suffixed with backslashes already. The final segment
// is stripped of its trailing backslash if it has one. Defaults to "/".
//
// NOTE: This is specifically for URL/URI paths, not platform-specific file paths.
func Url(segments ...string) string {
	return Path("/", segments...)
}

// Path combines the path segments using the provided delimiter. If the
// delimiter is empty, the default platform-specific delimiter will be used. If
// the path is empty after concatenation, the delimiter will be used as a default
// value.
func Path(delim string, segments ...string) string {
	if delim == "" {
		delim = DefaultPathDelim
	}
	var path string
	for i, segment := range segments {
		spaceTrimmed := strings.Trim(segment, " ")
		if i == 0 {
			if len(spaceTrimmed) == 0 {
				continue
			}
			if strings.HasPrefix(segment, delim) {
				path = fmt.Sprintf("%s%s", delim, strings.Trim(spaceTrimmed, delim))
			} else {
				path = strings.Trim(spaceTrimmed, delim)
			}
			continue
		}
		fullTrimmed := strings.Trim(segment, " "+delim)
		if len(fullTrimmed) == 0 && i+1 < len(segments) && strings.Trim(segment, " ") != delim {
			continue
		}
		path = fmt.Sprintf("%s%s%s", strings.TrimRight(path, delim), delim, fullTrimmed)
	}
	if len(path) == 0 {
		path = delim
	}
	return path
}
