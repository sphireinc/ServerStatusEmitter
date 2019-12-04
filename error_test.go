package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestLogErrorAndContinue(t *testing.T){

}

func TestHandleErrorFatal(t *testing.T){

}

func TestFormat(t *testing.T){

}

func TestTrace (t *testing.T) {
	tests := []struct {
		err      error
		expected string
	}{
		{errors.New("ErrorMessage"), "error_test.go <31> ServerStatusEmitter.TestTrace.func1(): ErrorMessage"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual := Trace(test.err)
			if actual != test.expected {
				t.Fatalf("expected '%s', got '%s'", test.expected, actual)
			}
		})
	}
}