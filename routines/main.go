package main

import (
	"fmt"
	"net/http"
)

func main() {
	/**links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.amazon.com",
		"http://www.microsoft.com",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLinkStatus(link, ch)
	}

	for l := range ch {
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLinkStatus(link, ch)
		}(l)
	}**/

	c := make(chan string)
	c <- "Hi there!"
}

func checkLinkStatus(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!")
		ch <- link
		return
	}
	fmt.Println(link, "is active")
	ch <- link
}
