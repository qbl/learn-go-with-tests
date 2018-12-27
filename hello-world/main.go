package main

import "fmt"

const helloPrefix = "Hello"

// HelloWorld prints "Hello World!"
func HelloWorld() string {
	return "Hello World!"
}

// Hello prints "Hello [name]!"
func Hello(name string, language string) string {
	if name == "" {
		return fmt.Sprintf("%s!", helloPrefix)
	}

	if language == "spanish" {
		return fmt.Sprintf("Hola %s!", name)
	}

	return fmt.Sprintf("%s %s!", helloPrefix, name)
}

func main() {
	fmt.Println(HelloWorld())
	fmt.Println(Hello("Iqbal", "spanish"))
}
