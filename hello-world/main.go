package main

import "fmt"

const helloPrefix = "Hello"

// HelloWorld prints "Hello World!"
func HelloWorld() string {
	return "Hello World!"
}

// Hello prints "Hello [name]!"
func Hello(name string) string {
	if name == "" {
		return fmt.Sprintf("%s!", helloPrefix)
	}
	return fmt.Sprintf("%s %s!", helloPrefix, name)
}

func main() {
	fmt.Println(HelloWorld())
	fmt.Println(Hello("Iqbal"))
}
