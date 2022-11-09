package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func printGreetings(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Helo!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
func main() {
	e := englishBot{}
	printGreetings(e)
	s := spanishBot{}
	printGreetings(s)
}
