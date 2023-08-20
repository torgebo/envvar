package envvar

import (
	"errors"
	"testing"
)

func CmpValue[V comparable](t *testing.T, ev EnvVar[V], expVar V) {
	if val := ev.Value(); val != expVar {
		t.Errorf("expected ev.Value()='%v', got '%v'", expVar, val)
	}
}

func CmpStringValue[V any](t *testing.T, ev EnvVar[V], expSVal string) {
	if sval := ev.StringValue(); sval != expSVal {
		t.Errorf("expected ev.StringValue()='%s', got '%s'", expSVal, sval)
	}
}

func CmpName[V any](t *testing.T, ev EnvVar[V], expName string) {
	if name := ev.Name(); name != expName {
		t.Errorf("expected ev.Name()='%s', got '%s'", expName, name)
	}
}

func CmpDescription[V any](t *testing.T, ev EnvVar[V], expDescription string) {
	if desc := ev.Description(); desc != expDescription {
		t.Errorf("expected ev.Description()='%s', got '%s'", expDescription, desc)
	}
}

func CmpString[V any](t *testing.T, ev EnvVar[V], expString string) {
	if str := ev.String(); str != expString {
		t.Errorf("expected ev.String()='%s', got '%s'", expString, str)
	}
}

// EnvVarTest runs test by
//  1. mocking out the underlying `osReader` function
//  2. Calling ev.Set()
//     a. comparing expReadErr with actual output of `ev.Set()`
//     b. observing that the call to `osReader` is performed
//  3. Consecutively comparing the values provided by the
//     other interface methods.
func EnvVarTest[V comparable](
	t *testing.T,
	ev EnvVar[V],
	mockOsReader func(name string) (string, bool),
	expReadErr error,
	expName string,
	expDescription string,
	expSVal string,
	expVar V,
	expString string,
) {
	var observeRead bool
	OsReader = func(name string) (string, bool) {
		observeRead = true
		return mockOsReader(name)
	}
	if err := ev.Set(); err != nil {
		if expReadErr != nil {
			if errors.Is(err, expReadErr) {
				return
			}
			t.Errorf("expected expReadErr='%s', got '%s'", expReadErr, err)
			return
		}
		t.Errorf("got unexpected error: %s", err)
		return
	}

	if !observeRead {
		t.Errorf("unable to observe any calls to mocked osReader")
		return
	}

	CmpName(t, ev, expName)
	CmpDescription(t, ev, expDescription)
	CmpStringValue(t, ev, expSVal)
	CmpValue(t, ev, expVar)
	CmpString(t, ev, expString)
}
