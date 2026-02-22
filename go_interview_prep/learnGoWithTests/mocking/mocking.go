package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countdownStart = 3
	finalWord      = "GO!"
)

type Sleeper interface {
	Sleep()
}

type RealSleeper struct{}

func (d *RealSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func CountDown(w io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		_, err := fmt.Fprintln(w, i)
		if err != nil {
			return
		}
		sleeper.Sleep()
	}

	_, err := fmt.Fprint(w, finalWord)
	if err != nil {
		return
	}
}
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
func main() {
	realSleeper := &RealSleeper{}
	CountDown(os.Stdout, realSleeper)
}
