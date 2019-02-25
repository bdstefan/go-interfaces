package main

import "fmt"

type bot interface {
	getGreetings() string
}

type englishBot struct {
	message string
}
type spanishBot struct {
	message string
}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}

func (eb englishBot) getGreetings() string  {
	return eb.message
}

func (sb spanishBot) getGreetings() string  {
	return sb.message
}
