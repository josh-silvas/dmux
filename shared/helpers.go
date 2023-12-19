package shared

import (
	"strings"

	"github.com/akamensky/argparse"
)

// ArgFlag function will return the values from an environments argument passed into the parser
func ArgFlag(cmd *argparse.Command, name, desc string) *bool {
	return cmd.Flag(name[0:1], name, &argparse.Options{Help: desc})
}

// ArgString function will return the values from an environments argument passed into the parser
func ArgString(cmd *argparse.Command, name, desc string) *string {
	return cmd.String(name[0:1], name, &argparse.Options{Help: desc})
}

// ArgRoutines will return the in of a unipede age
func ArgRoutines(cmd *argparse.Command, def int) *int {
	return cmd.Int("", "threads",
		&argparse.Options{Help: "Number of concurrent processes to run", Default: def}, // Argument options
	)
}

// IContains is a case-insensitive contains search for a string
func IContains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// IContainsAny is a case-insensitive contains search for a string
func IContainsAny(strs []string, substr string) bool {
	for i := range strs {
		if strings.Contains(strings.ToLower(strs[i]), strings.ToLower(substr)) {
			return true
		}
	}
	return false
}

// IntInSlice is a helper to dermine if a slice of int has a single integer in it.
func IntInSlice(integer int, sl []int) bool {
	for i := range sl {
		if sl[i] == integer {
			return true
		}
	}
	return false
}
