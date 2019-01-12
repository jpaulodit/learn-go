package main

import "fmt"

// This example is using Type Assertion to make sure that the error returned in the expected
// custom error type.

type imgproxyError struct {
	StatusCode    int
	Message       string
	PublicMessage string
}

func (e *imgproxyError) Error() string {
	return e.Message
}

func newError(status int, msg string, pub string) *imgproxyError {
	return &imgproxyError{
		StatusCode:    status,
		Message:       msg,
		PublicMessage: pub,
	}
}

func reportError(e *imgproxyError) {
	fmt.Println(e)
}

func main() {
	defer func() {
		if rerr := recover(); rerr != nil {
			if err, ok := rerr.(error); ok {
				if ierr, ok := err.(*imgproxyError); ok {
					reportError(ierr)
				}
			}
		}
	}()
	panic(newError(500, "Unexpected error", "Internal error"))
}

/* ----------------------- */

// This example shows how you can create custom error message that you can use at certain packages.

type ValidationErr struct {
	// We will store validation errors here.
	// The key is a field name, and the value is a slice of validation messages.
	ErrorMessages map[string][]string
}

func FormatErrors(e map[string][]string) string {
	return "Some formatted error from e"
}

func (e *ValidationErr) Error() string {
	return FormatErrors(e.ErrorMessages)
}


