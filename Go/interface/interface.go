package main

import "fmt"

 type bot interface {
	getGreeting() string
 }

type englishBot struct{}
type spanishBot struct{}

type logWriter struct{}


func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb) 

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return leng(bs), nil
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}


