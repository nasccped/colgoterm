package colored

import "fmt"

const (
	escapePrefix string = "\x1b["
	escapeReset  string = "\x1b[0m"
)

// Generates a normal foreground color escape over an int.
func nfg(val int) string {
	return genEscape(30 + val)
}

// Generates a bright foreground color escape over an int.
func bfg(val int) string {
	return genEscape(90 + val)
}

// Generates a normal background color escape over an int.
func nbg(val int) string {
	return genEscape(40 + val)
}

// Generates a bright background color escape over an int.
func bbg(val int) string {
	return genEscape(100 + val)
}

// Generates a new colored escape from a given int value.
func genEscape(val int) string {
	return fmt.Sprintf("%s%dm", escapePrefix, val)
}
