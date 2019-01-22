package main

import "fmt"
import "strconv"

func deferred_func_args_evaluated_early() {
	i := 0
	defer fmt.Println("deferred function argumented evaluted early: " + strconv.Itoa(i))
	i++
	return
}

func deferred_func_executed_LIFO() {
	fmt.Println("deferred function executed LIFO")
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func main() {
	deferred_func_args_evaluated_early()
	deferred_func_executed_LIFO()
}
