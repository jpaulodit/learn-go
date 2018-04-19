package main

import ("fmt"
        "time"
)

func greeting() string  {
    return "Hello the time is now " + time.Now().String()
}

func main() {
    fmt.Println(greeting())

    // example of declaring variable then assignment
    var now string
    now = time.Now().String()
    fmt.Println(now)
}
