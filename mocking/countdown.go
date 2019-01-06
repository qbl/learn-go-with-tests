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

// DefaultSleeper implements Sleeper
type DefaultSleeper struct{}

// Sleep implements time.Sleep()
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
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
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
