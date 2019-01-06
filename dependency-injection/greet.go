package main

import (
	"fmt"
	"io"
	"os"
)

// Greet accepts pointer to writer buffer and a name, and prints "Hello, <name>!" to writer buffer
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s!", name)
}

func main() {
	Greet(os.Stdout, "Iqbal")
}
