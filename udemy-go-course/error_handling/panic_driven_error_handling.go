package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/pkg/errors"
)

func HandleError(e error) {
	fmt.Println("Error is", e)
}

// **** Golang idiomatic error handling strategy with multiple "if err != nil" checks

// The parent function
func LogSqrt(str string) {
	f, err := StringToSqrt(str)
	/** 1 **/
	if err != nil {
		HandleError(err)
		return
	}
	log.Printf("Sqrt of %s is %f", str, f)
}

// Checks that input can be converted to a float
func StringToSqrt(str string) (float64, error) {
	f, err := strconv.ParseFloat(str, 64)
	/** 2 **/
	if err != nil {
		return 0, err
	}

	f, err = Sqrt(f)
	/** 3 **/
	if err != nil {
		return 0, err
	}

	return f, nil
}

// Performs the actual square root
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("cannot calculate sqrt of a negative number")
	} else {
		return math.Sqrt(f), nil
	}
}

// *** A panic driven approach for error handling. The idea is simple, "panic" whenever a nested
// *** invocation returns an error, and then recover from it in a single dedicated place.
// *** Places where "return 0, err" have been replaced with checkErr() calls.

// A simple helper to panic on errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func LogSqrtPanic(str string) {
	// recover() pairs with panic() and the method is only meaningful inside a deferred function.
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				HandleError(err)
			} else {
				// Recover returned something that is not an error, so "re-panic"
				panic(r)
			}
		}
	}()
	log.Printf("sqrt of %s is %f", str, StringToSqrtPanic(str))
}

func StringToSqrtPanic(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	checkErr(err)

	f, err = Sqrt(f)
	checkErr(err)

	return f
}

// *** Panic driven error handling approach that uses 3rd party errors package to provide more
// *** context to the errors. We are replacing the standard errors package with a 3rd party from
// *** github.com/pkg/errors

func checkErr2(err error, msg string) {
	if err != nil {
		panic(errors.WithMessage(err, msg))
	}
}

func checkErrWithStack(err error, msg string) {
	if err != nil {
		panic(errors.Wrap(err, msg))
	}
}

func LogSqrtPanic2(str string) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				log.Printf("%+v", err)
				HandleError(err)
			} else {
				panic(r)
			}
		}
	}()
	log.Printf("sqrt of %s is %f", str, StringToSqrtPanic2(str))
}

func StringToSqrtPanic2(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	checkErrWithStack(err, "Failed to parse")

	f, err = Sqrt(f)
	checkErr2(err, "Failed to sqrt")

	return f
}

func main() {
	LogSqrt("4")
	LogSqrtPanic("16")
	LogSqrtPanic("-10")
	LogSqrtPanic2("-1")
	LogSqrtPanic2("-1abcd")
}
