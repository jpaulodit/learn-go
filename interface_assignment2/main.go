package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f := os.Args[1]
	fmt.Println("File to open is:", f)

	my_file, err := os.Open(f)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, my_file)
}
