package main

import (
	"fmt"
	"io"
	"os"
)

/**
* Program to read file & print its contents on standard output
 */
func main() {
	arguments := os.Args
	filepath := arguments[1]
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
