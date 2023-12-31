// Package appvar provides the appVar methods that
package appvar

import (
	"errors"
	"fmt"
	"strings"

	"github.com/torgebo/envvar"
)

var (
	ErrNotSet = errors.New("variable has not been set")
	errPanic  = func(e error) {
		panic(e)
	}
)

// NewStringed creates a envvar.EnvVar[string] that reads in a string
func NewStringed(varname, description string) envvar.EnvVar[string] {
	parser := func(strvalue string) (string, error) {
		return strvalue, nil
	}
	return New[string](varname, description, parser)
}

// New creates a envvar.EnvVar with a custom parser of type `T`
func New[T any](varname, description string, parser func(string) (T, error)) envvar.EnvVar[T] {
	if err := validate[T](varname, description, parser); err != nil {
		errPanic(fmt.Errorf("exiting: envvar: %w", err))
	}
	var t T
	return &appVar[T]{
		varname:     varname,
		description: description,
		stringvalue: "",
		varvalue:    t,
		setcalled:   false,
		valuecalled: false,
		parser:      parser,
	}
}

func validate[T any](varname, description string, parser func(string) (T, error)) error {
	if strings.TrimSpace(varname) == "" {
		return fmt.Errorf("invalid varname '%s'", varname)
	}
	if strings.TrimSpace(description) == "" {
		return fmt.Errorf("invalid description '%s'", description)
	}
	if parser == nil {
		return fmt.Errorf("nil parser func")
	}
	return nil
}

// appVar provides environment variable `varname`
type appVar[T any] struct {
	varname     string
	description string
	stringvalue string
	varvalue    T
	setcalled   bool
	valuecalled bool
	parser      func(string) (T, error)
}

func (av *appVar[T]) Name() string {
	return av.varname
}

func (av *appVar[T]) Description() string {
	return av.description
}

func (av *appVar[T]) Set() error {
	if av.setcalled {
		return nil
	}
	name := av.Name()
	strval, found := envvar.OsReader(name)
	if !found {
		return fmt.Errorf("no such value '%s'", name)
	}
	av.stringvalue = strval
	value, err := av.parser(strval)
	if err != nil {
		return err
	}
	av.varvalue = value
	av.setcalled = true
	return nil
}

func (av *appVar[T]) StringValue() string {
	return av.stringvalue
}

func (av *appVar[T]) Value() T {
	if !av.setcalled {
		errPanic(fmt.Errorf("exiting: envvar: %w", ErrNotSet))
	}
	av.valuecalled = true
	return av.varvalue
}

func (av *appVar[T]) ValueRead() error {
	if !av.valuecalled {
		return fmt.Errorf(
			"envvar: '%s': %w",
			av.varname, envvar.ErrValueNotRead,
		)
	}
	return nil
}

func (av *appVar[T]) String() string {
	return fmt.Sprintf("%s: %s", av.Name(), av.Description())
}
