// envvar provides the interface that environment variables should implement
// as well as the `OsReader` method for wrapping os.LookupEnv
package envvar

import (
	"os"
	"errors"
)

// OsReader is the preferred method for reading system environment variables
// as it will allow for running of our tooling, like tests
var OsReader = func(name string) (string, bool) {
	return os.LookupEnv(name)
}

var ErrValueNotRead = errors.New("Value not read")

// EnvVar is an entity holding standard information on an environment variable
type EnvVar[T any] interface {
	// Name provides the entity name
	Name() string

	// Description provides the entity description.
	Description() string

	// Set reads, parses and sets Value
	Set() error

	// StringValue should be called after Set to
	// return the unparsed environment variable value
	//
	// Prefer Value() over StringValue() when possible.
	StringValue() string

	// Value should be called after Set to
	// return the parsed environment variable value
	//
	// If Set() has not been called, Value must panic.
	Value() T

	// ValueRead should be called after a call to Value
	//
	// Use this to check whether Value() has been called.
	// If Value has not previously been called it will return
	// an error that Is of type `ErrValueNotRead`.
	ValueRead() error

	// String provides the full variable documentation, including
	// Name(), Description() and StringValue() if Set() has been
	// successfully called
	String() string
}
