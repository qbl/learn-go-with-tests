package main

import (
	"fmt"
	"io"
	"os"
)

// Countdown accepts output buffer, and prints it
func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}

func main() {
	Countdown(os.Stdout)
}
