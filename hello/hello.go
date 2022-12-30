package main

import "fmt"

const portuguese = "Portuguese"
const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "

// Hello sends a greeting string
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == portuguese {
		return portugueseHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
