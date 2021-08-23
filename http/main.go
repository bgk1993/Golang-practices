package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//fmt.Println(response)

	//bs := make([]byte, 9999)
	//response.Body.Read(bs)
	//fmt.Println(string(bs))

	io.Copy(os.Stdout, response.Body)
}
