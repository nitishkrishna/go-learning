package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// worker thread, reads from jobs, writes to results channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func workerSynced(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Synced worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("Synced worker", id, "finished job", j)
		results <- j * 2
	}
	wg.Done()
}

func workerPool() {

	// Job queue and results queue
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Spawns 2 worker routines
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Spawn 15 jobs and Close channel to indicate no more jobs
	for j := 1; j <= 15; j++ {
		jobs <- j
	}
	close(jobs)

	// Read all results
	for a := 1; a <= 15; a++ {
		fmt.Println(<-results)
	}
}

var wg sync.WaitGroup

func workerPoolResCombined() {

	// Job queue and results queue
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Spawns 2 worker routines
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go workerSynced(w, jobs, results)
	}

	// Spawn 15 jobs and Close channel to indicate no more jobs
	for j := 1; j <= 15; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	close(results)

	// Read all results
	for a := range results {
		fmt.Println("Result", a)
	}
}

// Rate Limiting Request Serving using tickers

func eventGenerator(nreqs int, c chan<- time.Time) {
	for i := 1; i <= nreqs; i++ {
		c <- time.Now()
	}
}

func rateLimiting() {
	requests := make(chan time.Time, 5)
	eventGenerator(5, requests)
	close(requests)

	// Use limiter to time how to drain requests
	limiter := time.Tick(1 * time.Second)
	for req := range requests {
		<-limiter // Blocks for 1 s
		fmt.Println("request generated at ", req)
		fmt.Println("request drained at ", time.Now())
	}

	// Generate bursty Limiter to allow burst of N reqs before rate Limiting
	burstyLimiter := make(chan time.Time, 3)
	eventGenerator(3, burstyLimiter)

	// Allow to add events to above limiter using Ticker
	go func() {
		for t := range time.Tick(1 * time.Second) {
			burstyLimiter <- t
			// Can add N events here to allow burts of N every second
		}
	}()

	// Generate more requests
	burstyReqs := make(chan time.Time, 5)
	eventGenerator(5, burstyReqs)
	close(burstyReqs)
	for req := range burstyReqs {
		<-burstyLimiter // Initially allows 3 request burst before being limited
		fmt.Println("request generated at ", req)
		fmt.Println("request drained at ", time.Now())
	}

}

// Atomic Counters - Unsigned ints accessed by all go routines
func atomicCounters() {

	// counter
	var ops uint64

	// 50 Simultaneous go routines that increment the counter
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// pass counter address
				atomic.AddUint64(&ops, 1)

				// Wait a bit between increments.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Wait a second to allow some ops to accumulate.
	time.Sleep(2 * time.Second)

	// extract copy of counter value
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("Number of operations in 2 seconds:", opsFinal)
}

func shareStateViaMutexes() {
	// Maintain state in map
	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	// Counters to count ops
	var readOps uint64
	var writeOps uint64

	// 100 Go Routines to read state from same var
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				// Read random key and increment ops count
				key := rand.Intn(5)
				// Locking allows safe access - Blocking call
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				// Wait a bit between reads.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 10 Go Routines to simulate writes
	for w := 0; w < 10; w++ {
		go func() {
			for {
				// Write random Key in map and increment ops
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Allow two seconds of operations
	time.Sleep(2 * time.Second)

	// report final operation counts.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// Show final state
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}

// Read state req
type readOp struct {
	key  int
	resp chan int
}

// Write state req
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func stateManager(reads <-chan *readOp, writes <-chan *writeOp) {
	var state = make(map[int]int)
	for {
		select {
		case read := <-reads:
			read.resp <- state[read.key]
		case write := <-writes:
			state[write.key] = write.val
			write.resp <- true
		}
	}
}

func shareStateViaGoRoutines() {

	// operation counts
	var readOps uint64
	var writeOps uint64

	// Channels to issue read/write requests
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// Spawn State Owner Routine
	go stateManager(reads, writes)

	// 100 Go routines to spawn readOps
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 10 Go routines to spawn writeOps
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(2 * time.Second)

	// Capture final Op Counters
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

}

func main() {
	workerPool()
	workerPoolResCombined()
	rateLimiting()
	atomicCounters()
	shareStateViaMutexes()
	shareStateViaGoRoutines()
}
