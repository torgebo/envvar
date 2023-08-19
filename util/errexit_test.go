package util

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

type TestErrCase struct {
	ErrInput   error
	OutString  string
	ExitCalled bool
}

var ErrExitCases = []struct {
	ErrInput   error
	OutString  string
	ExitCalled bool
}{
	{
		ErrInput:   errors.New("error here"),
		OutString:  "exiting: error here",
		ExitCalled: true,
	},
	{
		ErrInput:   nil,
		OutString:  "",
		ExitCalled: false,
	},
}

func TestErrExit(t *testing.T) {
	for ind, testcase := range ErrExitCases {
		t.Run(
			fmt.Sprintf("case: %d, error '%s'", ind, testcase.ErrInput),
			func(t *testing.T) {
				var b bytes.Buffer
				stdErr = &b
				var called = false
				osExit = func(status int) {
					called = true
				}
				ErrExit(testcase.ErrInput)
				if got, exp := called, testcase.ExitCalled; got != exp {
					t.Errorf("expected os.Exit to be called=%t, got %t", exp, got)
				}
				if got, exp := b.String(), testcase.OutString; got != exp {
					t.Errorf("expected output '%s', got '%s'", exp, got)
				}
			},
		)

	}

}
