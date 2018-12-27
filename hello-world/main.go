package main

import "fmt"

const french = "French"
const frenchHelloPrefix = "Bonjour"
const helloPrefix = "Hello"
const spanish = "Spanish"
const spanishHelloPrefix = "Hola"

// HelloWorld prints "Hello World!"
func HelloWorld() string {
	return "Hello World!"
}

// Hello prints "Hello [name]!"
func Hello(name string, language string) string {
	if name == "" {
		return fmt.Sprintf("%s!", helloPrefix)
	}

	if language == "Spanish" {
		return fmt.Sprintf("%s %s!", spanishHelloPrefix, name)
	}

	if language == "French" {
		return fmt.Sprintf("%s %s!", frenchHelloPrefix, name)
	}

	return fmt.Sprintf("%s %s!", helloPrefix, name)
}

func main() {
	fmt.Println(HelloWorld())
	fmt.Println(Hello("Iqbal", "Spanish"))
}
