package main

import (
	"bytes"
	"fmt"
)

// Countdown accepts output buffer, and prints it
func Countdown(out *bytes.Buffer) {
	fmt.Fprint(out, "3")
}

func main() {}
