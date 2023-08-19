// Package util provides some tool that comes in handy
package util

// Must panics on error
//
// Users are encourage to instead use ErrExit when possible.
func Must(e error) {
	if e != nil {
		panic(e)
	}
}
