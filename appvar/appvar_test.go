package appvar

import (
	"fmt"
	"github.com/torgebo/envvar"
	"strconv"
	"testing"
	"time"
)

var number int = 3
var nums string = fmt.Sprintf("%d", number)

var testTime = time.Now().Truncate(time.Second)
var testTimeString = testTime.Format(time.RFC3339)

func osReader1(name string) (string, bool) {
	if name == `testtimevar` {
		return testTimeString, true
	} else if name == `testintvar` {
		return nums, true
	}
	return "", false
}

var timeParser = func(timestring string) (time.Time, error) {
	return time.Parse(time.RFC3339, timestring)
}

var intParser = func(numstring string) (int, error) {
	return strconv.Atoi(numstring)
}

func TestAppVarTime(t *testing.T) {
	ev := New(`testtimevar`, `this should be a datetime`, timeParser)

	envvar.EnvVarTest(
		t,
		ev,
		osReader1,
		nil,
		`testtimevar`,
		`this should be a datetime`,
		testTime.Format(time.RFC3339),
		testTime,
		`testtimevar: this should be a datetime`,
	)

}

func TestEnvVar(t *testing.T) {
	envvar.OsReader = osReader1

	ev1 := New(`testtimevar`, `this should be a datetime`, timeParser)
	ev2 := New(`testintvar`, `This should be an integer`, intParser)
	vvs := envvar.ToVars(ev1, ev2)

	if err := vvs.Set(); err != nil {
		t.Error(err)
		return
	}

	if gotTime := ev1.Value(); gotTime != testTime {
		t.Errorf("expected ev1.Value()=%v, got %v", testTime, gotTime)
		return
	}
	if gotNum := ev2.Value(); gotNum != number {
		t.Errorf("expected ev2.Value()=%d, got %d", number, gotNum)
		return
	}

	description := vvs.String()
	if exp := `Required Environment Variables:
testtimevar: this should be a datetime
testintvar: This should be an integer
`; exp != description {
		t.Errorf("expected vvs.String()='%s', got '%s'", exp, description)
	}
}

func TestEnvVarValueBeforeSet(t *testing.T) {
	var errPanicCalled = false
	errPanic = func(_ error) {
		errPanicCalled = true
	}
	ev := New(`testtimevar`, `this should be a datetime`, timeParser)
	ev.Value()
	if !errPanicCalled {
		t.Errorf("expected errPanic to be called, got errPanicCalled=%t", errPanicCalled)
	}

}
