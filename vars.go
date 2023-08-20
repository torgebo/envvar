package envvar

import (
	"fmt"
	"strings"
)

var (
	errPanic = func(e error) {
		panic(e)
	}
)

type SetString interface {
	Set() error
	String() string
}

// ToVars is a convenient way to construct Vars
func ToVars(evs ...SetString) Vars {
	return Vars(evs)
}

// Vars is handy to Set in and output a summary string
// from several EnvVar-s
type Vars []SetString

// Set reads in all variables in the collection
func (vs Vars) Set() error {
	var err error
	for _, ev := range vs {
		if err = ev.Set(); err != nil {
			return err
		}
	}
	return nil
}

// String formats a helper description string from all variables in the collection
func (vs Vars) String() string {
	var builder strings.Builder
	if _, err := builder.WriteString("Required Environment Variables:\n"); err != nil {
		errPanic(fmt.Errorf("exiting: envvar: %w", err))
	}
	for _, ev := range vs {
		if _, err := builder.WriteString(ev.String()); err != nil {
			errPanic(fmt.Errorf("exiting: envvar: %w", err))
		}
		if _, err := builder.WriteString("\n"); err != nil {
			errPanic(fmt.Errorf("exiting: envvar: %w", err))
		}
	}
	return builder.String()
}
