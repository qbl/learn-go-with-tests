package main

import (
	"fmt"
	"io"
	"os"
)

// Countdown accepts output buffer, and prints it
func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
