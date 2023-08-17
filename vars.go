package envvar

import "strings"

type ReadString interface {
	Read() error
	String() string
}

// ToVars is a convenient way to construct Vars
func ToVars(evs ...ReadString) Vars {
	return Vars(evs)
}

// Vars is handy to Read in and output a summary string
// from several EnvVar-s
type Vars []ReadString

// Read reads in all variables in the collection
func (vs Vars) Read() error {
	var err error
	for _, ev := range vs {
		if err = ev.Read(); err != nil {
			return err
		}
	}
	return nil
}

// String formats a helper description string from all variables in the collection
func (vs Vars) String() string {
	var builder strings.Builder
	if _, err := builder.WriteString("Environment Variables:\n"); err != nil {
		panic(err)
	}
	for _, ev := range vs {
		if _, err := builder.WriteString(ev.String()); err != nil {
			panic(err)
		}
		if _, err := builder.WriteString("\n"); err != nil {
			panic(err)
		}
	}
	return builder.String()
}
