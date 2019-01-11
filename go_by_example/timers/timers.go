package main

import (
	"fmt"
	"time"
)

func timers() {

	// Timers cause an even in a channel after specified time
	// This timer will wait 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)

	// Blocks till timer expires
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// Better than time.Sleep as timer can be cancelled in between if needed
	timer2 := time.NewTimer(5 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	time.Sleep(2 * time.Second)
	stop2 := timer2.Stop() // stopping the timer which hasn't expired
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

func tickers() {

	// Tickers cause events every x time interval
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		// The channel has multiple events so we need to range over them
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// Tickers can be stopped like timers.
	time.Sleep(3200 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func main() {
	timers()
	tickers()
}
