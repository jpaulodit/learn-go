package main

import "fmt"

func main()  {
    // way 1 to create map
    colors := map[string]string{
        "red": "ff0000",
        "blue": "00ff00",
        "green": "bbff00",
    }
    fmt.Println(colors)

    var colors2 map[string]string
    fmt.Println(colors2)

    // in this approach, you don't know the KV that you want yet
    colors3 := make(map[string]string)
    colors3["white"] = "ffffff"
    colors3["green"] = "blah"
    fmt.Println(colors3)

    delete(colors3, "green")
    fmt.Println(colors3)

    printMap(colors)
}

// This is how to loop over map
func printMap(c map[string]string)  {
    fmt.Println("\nExample code to print map.")
    for k, v := range c{
        fmt.Println("Hex code for", k, "is", v)
    }
}
