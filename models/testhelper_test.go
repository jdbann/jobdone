package models_test

import "regexp"

var ansiCodes = regexp.MustCompile(`\x1b\[[\d;]m`)

func stripAnsi(s string) string {
	return ansiCodes.ReplaceAllString(s, "")
}
