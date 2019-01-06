package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Greet accepts pointer to writer buffer and a name, and prints "Hello, <name>!" to writer buffer
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s!", name)
}

// MyGreeterHandler accepts http.ResponseWriter and pointer to http.Request, and prints greeting
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	Greet(os.Stdout, "Iqbal")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
