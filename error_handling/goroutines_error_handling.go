// This example doesn't really happen commonly because in this example, all errors are collected
// and processed.
package main

import (
	"errors"
	"fmt"
	"sync"
)


func ErrorHandler(e error) {
	fmt.Println("Error is", e)
}


func dangerousAction() error {
	return errors.New("some random error")
}

func main() {
	errCh := make(chan error, 2)

	var wg sync.WaitGroup
	// We will launch 2 goroutines.
	wg.Add(2)

	// Goroutine #1
	go func() {
		defer wg.Done()

		if err := dangerousAction(); err != nil {
			errCh <- err
			return
		}
	}()

	// Goroutine #2
	go func() {
		defer wg.Done()

		if err := dangerousAction(); err != nil {
			errCh <- err
			return
		}
	}()

	wg.Wait() // this blocks execution
	close(errCh)

	// Loop over channel and collect all errors.
	for err := range errCh {
		ErrorHandler(err)
	}
}
