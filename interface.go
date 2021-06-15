package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type SpanishBot struct{}

func main() {
	eb := englishBot{}
	sb := SpanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func (englishBot) getGreeting() string {
	return "Hi There"
}

func (SpanishBot) getGreeting() string {
	return "Holo"
}
