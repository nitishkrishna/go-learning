package main

import (
	"fmt"
	"sync"
	"time"
)

// Vid 19
// Need to add Go routines to WaitGroups
var waitGroup sync.WaitGroup
var chanWaitGroup sync.WaitGroup

func say(s string) {

	// Notify wait group that one Go Routine has completed
	// This function can ONLY be run as a Go Routine
	// Defer means this runs even if something else in func fails
	// This statement runs at end of say
	defer waitGroup.Done()

	defer cleanup()

	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		// Cause a panic to recover from
		if i == 2 {
			panic("Oh dear, a 2")
		}
	}

}

// Vid 20
func foo() {
	// This statement runs only when the rest of foo has run or has errored out
	defer fmt.Println("Function foo done")
	// Defers run in LIFO order, so below statement will be the first Defer to run
	defer fmt.Println("Is the function done?")

	fmt.Println("Running function foo")
}

// Vid 21
// Panic and Recover builtin functions
func cleanup() {
	if r := recover(); r != nil {
		// Recover from panic caused in say function
		// Note that cleanup function call MUST be added in function where panic is to be caught
		// Similar to try, except and catch
		fmt.Println("Recovered in cleanup func, panic reported:", r)
	}
}

// Vid 22 - Channels
func channelTest(c chan int, someVal int) {
	// Send some integer value to the channel
	c <- someVal * 5
}

// Vid 23 - Iterate Channels
func channelBuffTest(c chan int, someVal int) {
	// Collect values
	defer chanWaitGroup.Done()
	// Send some integer value to the channel
	c <- someVal * 10
}

func main() {
	// Vid 18
	// Go routine (lightweight thread)
	waitGroup.Add(1) // Sync
	go say("Hey")
	waitGroup.Add(1)
	go say("There")

	// Program should not finish before Go Routines finish
	// In Practice this is not a good idea
	// Need to explicitly sync them
	// One second time given for Routines to finish
	// time.Sleep(time.Second)

	// Vid 19
	waitGroup.Wait() // Wait for completion

	// Run a routine after the previous two routines
	waitGroup.Add(1)
	go say("Later")
	waitGroup.Wait()

	// Vid 20
	// Defer Statement
	// If our routine errors out before wg.Done(), wg.Wait() will wait forever
	// To prevent this, use Defer

	foo()

	// Vid 22
	// Channels - Send and receive values via Go Routines
	// Channel Operator is <-
	// Channel defined by Datatype and buffer size
	myChannelVal := make(chan int)
	go channelTest(myChannelVal, 10)
	go channelTest(myChannelVal, 17)

	// Read values from the channel
	v1 := <-myChannelVal
	v2 := <-myChannelVal

	fmt.Println(v1, v2)

	// Vid 23
	// Buffering and Iterating over Channels
	// Channel defined by Datatype and buffer size
	myBuffChannelVal := make(chan int, 10)

	for i := 0; i < 10; i++ {
		chanWaitGroup.Add(1)
		go channelBuffTest(myBuffChannelVal, i)
	}

	// Wait here for Channel to collect all values
	// Synchronise the routines
	chanWaitGroup.Wait()
	// Close the Channel
	close(myBuffChannelVal)

	// Iterate the channel size
	for item := range myBuffChannelVal {
		fmt.Println(item)
	}

}
