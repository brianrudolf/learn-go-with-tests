package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	germanHelloPrefix = "Guten tag, "
)

func greetingPrefix(language string) (prefix string) {
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
	return prefix
}

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

func main() {
	fmt.Println(Hello("Brian", ""))
}

