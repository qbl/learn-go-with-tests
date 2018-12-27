package main

import "fmt"

// HelloWorld prints "Hello World!"
func HelloWorld() string {
	return "Hello World!"
}

// Hello prints "Hello [name]!"
func Hello(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

func main() {
	fmt.Println(HelloWorld())
	fmt.Println(Hello("Iqbal"))
}
