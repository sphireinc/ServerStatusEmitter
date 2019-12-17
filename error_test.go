package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_LogErrorAndContinue(t *testing.T) {

}

func TestError_HandleErrorFatal(t *testing.T) {

}

func TestError_Format(t *testing.T) {
	tests := []struct {
		err      error
		expected string
	}{
		{errors.New("ErrorMessage"), "ERROR ErrorMessage"},
		{errors.New(""), "ERROR "},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual := format("Error", test.err)
			if actual != test.expected {
				t.Fatalf("expected '%s', got '%s'", test.expected, actual)
			}
		})
	}
}

func TestError_Trace(t *testing.T) {
	tests := []struct {
		err      error
		expected string
	}{
		{errors.New("ErrorMessage"), "error_test.go <46> ServerStatusEmitter.TestError_Trace.func1(): ErrorMessage"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual := trace(test.err)
			if actual != test.expected {
				t.Fatalf("expected '%s', got '%s'", test.expected, actual)
			}
		})
	}
}
