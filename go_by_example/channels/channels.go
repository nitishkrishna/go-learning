package main

import (
	"fmt"
	"time"
)

func sendAndReceive() {
	messages := make(chan string)
	go func() { messages <- "ping" }() // send
	msg := <-messages                  // receive
	fmt.Println(msg)
}

func bufferedChannel(size int) {
	messages := make(chan string, size) // channel size is 2
	messages <- "buffered"
	messages <- "channel"
	close(messages)
	for msg := range messages {
		fmt.Println(msg)
	}
}

// This is the function we'll run in a goroutine. The
// `done` channel will be used to notify another
// goroutine that this function's work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we're done.
	done <- true
}

func syncRoutinesUsingChannels() {

	// Start a worker goroutine, giving it the channel to
	// notify on.
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the
	// worker on the channel.
	<-done
}

// Send-only channel pings, exposing only one dir
func ping(pings chan<- string, msg string) {
	pings <- msg // Send to pings channel
	// Cannot Read
	// new_msg := <-pings --- this causes error below
	// invalid operation: <-pings (receive from send-only type chan<- string)
}

// Receive-only channel pings, exposing only one dir
// Send-only channel pongs
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings // Read from pings channel
	pongs <- msg   // Send to pongs channel
}

func channelDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func channelSelection() {

	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount
	// of time, to simulate e.g. blocking RPC operations
	// executing in concurrent goroutines.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Simultaneously wait on both channels for msgs
	for i := 0; i < 2; i++ { // Excpecting only two msgs
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func channelTimeouts() {
	// Prevents channels from waiting infinitely and blocking

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second) // Simulates some external call that returns
		c1 <- "result 1"
	}()

	select {
	case res := <-c1: // wait for result
		fmt.Println(res)
	case <-time.After(1 * time.Second): // timeout within 1 second if no result
		fmt.Println("timeout 1") // this case times out
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res) // we receive result within timeout
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

func nonBlockingChannels() {
	messages := make(chan string)
	signals := make(chan bool)

	// Non blocking receive
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Non blocking send
	// msg can't be send because no buffer size to store msg
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// Multichannel non-blocking receive
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func receiver(jobs chan int, done chan bool) {
	for {
		j, more := <-jobs
		if more {
			fmt.Println("received job", j)
		} else {
			fmt.Println("received all jobs")
			done <- true
			return
		}
	}
}

func closeChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	go receiver(jobs, done)
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
		time.Sleep(time.Second)
	}
	close(jobs) // Closing channel causes more to set to false and routine to end
	fmt.Println("sent all jobs")
	<-done // Wait via channel sync
}

func main() {
	//sendAndReceive()
	bufferedChannel(2)
	//syncRoutinesUsingChannels()
	//channelDirections()
	//channelSelection()
	//channelTimeouts()
	//nonBlockingChannels()
	closeChannel()
}
