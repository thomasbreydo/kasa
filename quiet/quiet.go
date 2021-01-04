package quiet

import (
	"os"
)

// SetupAndGetTeardown sets stdout and stderr to /dev/null, returning a function
// to reset them.
//
// The returned function should be immediately deferred.
func SetupAndGetTeardown() func() {
	null, _ := os.Open(os.DevNull)
	stdout, stderr := os.Stdout, os.Stderr
	os.Stdout, os.Stderr = null, null
	return func() {
		null.Close()
		os.Stdout, os.Stderr = stdout, stderr
	}
}
