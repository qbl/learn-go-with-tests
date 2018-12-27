package main

import "fmt"

const englishPrefix = "Hello"
const french = "French"
const frenchHelloPrefix = "Bonjour"
const spanish = "Spanish"
const spanishHelloPrefix = "Hola"

// HelloWorld prints "Hello World!"
func HelloWorld() string {
	return "Hello World!"
}

// Hello prints "Hello [name]!"
func Hello(name string, language string) string {
	if name == "" {
		return fmt.Sprintf("%s!", greetingPrefix(language))
	}

	return fmt.Sprintf("%s %s!", greetingPrefix(language), name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishPrefix
	}

	return prefix
}

func main() {
	fmt.Println(HelloWorld())
	fmt.Println(Hello("Iqbal", "Spanish"))
}
