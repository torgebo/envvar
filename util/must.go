// Package util provides some tool that comes in handy
package util

// Must panics on error
func Must(e error) {
	if e != nil {
		panic(e)
	}
}
