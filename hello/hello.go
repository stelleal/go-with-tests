package main

import "fmt"

const portuguese = "Portuguese"
const french = "French"
const italian = "Italian"
const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "
const frenchHelloPrefix = "Bonjour, "
const italianHelloPrefix = "Ciao, "

// Hello sends a greeting string
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case portuguese:
		prefix = portugueseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case italian:
		prefix = italianHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
