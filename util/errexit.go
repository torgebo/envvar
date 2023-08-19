package util

import (
	"fmt"
	"io"
	"os"
)

var (
	osExit           = os.Exit
	stdErr io.Writer = os.Stderr
)

// ErrExit prints error `e` and exits if `e` is non-nil
func ErrExit(e error) {
	if e == nil {
		return
	}
	err := fmt.Errorf("exiting: %w", e)
	fmt.Fprintf(stdErr, "%s", err)
	osExit(1)
}
