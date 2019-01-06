package main

import (
	"fmt"
	"io"
	"os"
)

const finalWord = "Go!"
const countdownStart = 3

// Countdown accepts output buffer, and prints it
func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
