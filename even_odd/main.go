package main

import "fmt"

func main() {

    // create
    slice_int := [11]int{}
    for index := range slice_int {
        slice_int[index] = index
    }

    for _, value := range slice_int {
        if value % 2 == 0 {
            fmt.Printf("%d is even\n", value)
        } else {
            fmt.Printf("%d is odd\n", value)
        }
    }
}