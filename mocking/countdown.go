package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

// Custom types
type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}

type SpyCountdownOperations struct {
	Calls []string
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

// Methods
func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Constants
const write = "write"
const sleep = "sleep"
const FinalWord = "Go!"
const CountdownStart = 3

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := CountdownStart; i > 0; i -= 1 {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}

	fmt.Fprintln(writer, FinalWord)
}

func main() {
	// sleeper := &DefaultSleeper{}
	sleeper := &ConfigurableSleeper{duration: 2 * time.Second, sleep: time.Sleep}
	Countdown(os.Stdout, sleeper)
}
