package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const write = "write"
const sleep = "sleep"

type SpySleeper struct {
	SleepCount int
}

func (s *SpySleeper) Sleep() {
	s.SleepCount++
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

type SpyCountDownOperations struct {
	Calls []string
}

func (s *SpyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountDown(t *testing.T) {
	t.Run("prints countdown from 3 to 1", func(t *testing.T) {
		buffer := new(bytes.Buffer)
		spySleeper := &SpySleeper{}

		CountDown(buffer, spySleeper)
		got := buffer.String()
		want := `3
2
1
GO!`
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		if spySleeper.SleepCount != countdownStart {
			t.Errorf("expected sleep to be called %v times, got %d", countdownStart, spySleeper.SleepCount)
		}
	})

	t.Run("prints countdown from 5 to 1 using a spy", func(t *testing.T) {
		spySleepPrinter := &SpyCountDownOperations{}
		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write}

		if !reflect.DeepEqual(spySleepPrinter.Calls, want) {
			t.Errorf("got %v want %v", spySleepPrinter.Calls, want)
		}
	})

}
func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.SetDurationSlept}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
