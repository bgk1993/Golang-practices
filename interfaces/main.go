package main

import "fmt"

type Bot interface {
	getGreetings() string
}

type EnglishBot struct{}

type SpanishBot struct{}

func main() {
	englishBot := EnglishBot{}
	spanishBot := SpanishBot{}

	printGreeting(englishBot)
	printGreeting(spanishBot)
}

func (e EnglishBot) getGreetings() string {
	return "Hi There!"
}

func (s SpanishBot) getGreetings() string {
	return "Hoila!"
}

func printGreeting(b Bot) {
	fmt.Println(b.getGreetings())
}
