package main

import "fmt"

type bot interface {
	getGreeting() string
}

// type englishBot struct{}
// type spanishBot struct{}
type roBot struct {
	language string
}

// func (englishBot) getGreeting() string {
// 	// Very custom logic for returning a greeting
// 	return "Hi there! How are you?"
// }

// func (spanishBot) getGreeting() string {
// 	// Very custom logic for returning a greeting
// 	return "¡Hola! ¿Cómo estás?"
// }

func (bot roBot) getGreeting() string {
	if bot.language == "english" {
		return "Hi there! How are you?"
	} else if bot.language == "spanish" {
		return "¡Hola! ¿Cómo estás?"
	}
	return ""
}

// func printGreeting(en englishBot) {
// 	fmt.Println(en.getGreeting())
// }

// func printGreeting(sp spanishBot) {
// 	fmt.Println(sp.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
