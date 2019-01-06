package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Sleeper is an interface to function Sleep()
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper has duration and sleep
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep implements time.Sleep
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Countdown accepts output buffer and sleeper interface, and prints the output buffer
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{5 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
