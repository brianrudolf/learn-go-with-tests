package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const germanHelloPrefix = "Guten tag, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	var prefix string

	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "English":
		prefix = englishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	case "German":
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	message := prefix + name
	return message
}

func main() {
	fmt.Println(Hello("Brian", ""))
}

