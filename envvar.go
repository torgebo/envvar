// envvar provides the interface that environment variables should implement
// as well as the `OsReader` method for wrapping os.LookupEnv
package envvar

import "os"

// OsReader is the preferred method for reading system environment variables
// as it will allow for running of our tooling, like tests
var OsReader = func(name string) (string, bool) {
	return os.LookupEnv(name)
}

// EnvVar is an entity holding standard information on an environment variable
type EnvVar[T any] interface {
	// Name provides the entity name
	Name() string

	// Description provides the entity description.
	Description() string

	// Read reads, parses and sets Value
	Read() error

	// StringValue should be called after Read to
	// return the unparsed environment variable value
	StringValue() string

	// Value should be called after Read to
	// return the parsed environment variable value
	Value() T

	// String provides the full variable documentation, including
	// Name(), Description() and StringValue() if Read() has been
	// successfully called
	String() string
}
