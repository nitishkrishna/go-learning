package main

import (
	"errors"
	"fmt"
	"math"
	"runtime"
	"strings"
	"sync"
	"time"

	"golang.org/x/tour/pic"
)

type person struct {
	name string
	age  int
}

// Simple two argument function
func mySum(x int, y int) int {
	return x + y
}

// Funcs can return errors
// Go doesn't have exceptions
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Cannot return SQRT of negative number")
	}

	return math.Sqrt(x), nil
}

func inc(x int) {
	x++
}

func incPointer(x *int) {
	*x++
}

// If can evaluate before condition
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// Simple Word count using Maps

func wordCount(s string) map[string]int {
	myMap := make(map[string]int)
	for _, x := range strings.Fields(s) {
		_, check := myMap[x]
		if check {
			myMap[x]++
		} else {
			myMap[x] = 1
		}
	}
	return myMap
}

// Passing function as argument to function

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Function closures
// Closure function can access vars outside the Body
// Here adder func returns a closure func that accesses Sum

func adder() func(int) int {
	sum := 0
	// Anonymous "closure" function returned
	return func(x int) int {
		sum += x
		return sum
	}
}

// Fibonacci closures
func fibonacci() func() int {
	x, y := 0, 0
	return func() int {
		if x == 0 && y == 0 {
			y++
			return 0
		}

		x = x + y
		y = x - y
		return x
	}
}

// Method Interfaces
// An object of Abser interface type can be assigned to any struct that has an Abs method
type abser interface {
	Abs() float64
}

// Example:
type myFloat struct {
	val float64
}

func (f *myFloat) Abs() float64 {
	if f == nil {
		fmt.Println("<nil>")
		return 0
	}
	if f.val < 0 {
		return float64(-f.val)
	}
	return float64(f.val)
}

func dataStructs() {
	// Array declare and assign vals
	fmt.Println("Arrays ")
	var arr [5]int
	arr[2] = 10
	arr[0] = 11
	fmt.Println(arr)

	b := [3]int{9, 8, 7}
	fmt.Println(b)

	// Dynamic Arrays - Slices - No fixed size

	fmt.Println("Slices ")

	slc := []int{1, 2, 3}
	slc = append(slc, 4)       // append a single element
	slc = append(slc, 5, 6, 7) // append multiple elements
	secondSlc := []int{8, 9, 10}
	slc = append(slc, secondSlc...) // append a slice
	fmt.Println(slc)
	slc = slc[1:4] // Subslice
	fmt.Println(slc)

	// Slices are references to Arrays

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	namesA := names[0:2]
	namesB := names[1:3]
	fmt.Println(namesA, namesB)

	namesB[0] = "XXX" // This changes names[1] as well
	fmt.Println(namesA, namesB)
	fmt.Println(names)

	// Map for key/value pairs

	fmt.Println("Maps ")

	myMap := make(map[string]int)
	myMap["Nitish"] = 22
	myMap["Alice"] = 83
	myMap["Bob"] = 100

	delete(myMap, "Bob")

	fmt.Println(myMap)
	// Check if an element is present in map
	elem, check := myMap["Eleven"]
	fmt.Println("The value:", elem, "Present?", check)

	// Creating struct of a type
	fmt.Println("Structs")
	p := person{name: "Nitish", age: 45}
	fmt.Println(p)
	fmt.Println("Age:", p.age)
	// Pointers to structs
	// More efficient for larger structs
	pPtr := &p
	personPtr := &person{"Alice", 55}
	// Doesn't need *
	pPtr.age = 50
	fmt.Println(p, personPtr)

}

func programmingConstructs() {
	// If condition
	x := 5
	fmt.Println("If cond ")
	if x > 6 {
		fmt.Println("More than 6")
	} else if x == 5 {
		fmt.Println("Equal to 5")
	} else {
		fmt.Println("Less than 6, Not 5")
	}

	// Loop - Only for loop in go

	fmt.Println("Loops ")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While
	fmt.Println("While Loop ")

	j := 7

	for j < 10 {
		fmt.Println(j)
		j++
	}

	// Iterate arr
	fmt.Println("Iterate Arr and Map ")

	newnames := []string{"Tom", "Dick", "Harry"}

	for idx, val := range newnames {
		fmt.Println("index:", idx, "value:", val)
	}

	// Iterate map
	capMap := make(map[string]string)
	capMap["India"] = "Delhi"
	capMap["Canada"] = "Ottowa"

	for key, val := range capMap {
		fmt.Println("key:", key, "value:", val)
	}

	// Switch Case
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	// Long If Else chains can be written as below:
	t := 100
	switch {
	case t < 12:
		fmt.Println("Good morning!")
	case t < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// If can evaluate before condition
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

}

func functionsAndPointers() {
	// Call function
	// lowerCase for internal functions
	fmt.Println("Functions ")

	r := mySum(2, 3)
	fmt.Println("Sum:", r)

	res, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sqrt:", res)
	}

	res, err = sqrt(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sqrt:", res)
	}

	// Variable memory address
	fmt.Println("Mem Addr and Pointers")
	i := 7
	fmt.Println(&i)

	// Value doesn't change as copy is made
	inc(i)
	fmt.Println(i)

	// Value changes, pointer passed
	incPointer(&i)
	fmt.Println(i)

	// Functions as values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	// Pass function as variable to function
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// Anonymous functions - no names

	func(msg string) {
		fmt.Println(msg)
	}("Anonymous Func Call")

	// Returning Anonymous functions
	fmt.Println("Closures")
	// Functions are bound to variables
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),    // the Pos function maintains its own sum var
			neg(-2*i), // same for neg function
		)
	}

	fmt.Println("Fibonacci closure")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	fmt.Println("Interfaces")

	var abserobj abser // var of interface type, NOT myfloat type
	myfloat := &myFloat{-math.Sqrt2}
	abserobj = myfloat
	fmt.Println(abserobj.Abs()) // myFloat's Abs method gets called
	// Interface can itself point to a nil object
	var mf *myFloat
	abserobj = mf

	fmt.Println("Interface type assertions") // To know what Type is your interface
	var intf interface{} = "hello"

	intfValue, checkConcreteType := intf.(string)
	fmt.Println(intfValue, checkConcreteType)

	intfValue2, checkConcreteType := intf.(float64)
	fmt.Println(intfValue2, checkConcreteType)

}

// Slice Exercise

// Pic ...
func Pic(dx, dy int) [][]uint8 {
	matrix := make([][]uint8, dy)

	for k := range matrix {
		matrix[k] = make([]uint8, dx)
	}

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			matrix[i][j] = uint8((i + j) / 2)
		}
	}
	return matrix
}

// channels
func chanSum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacciChannel(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select { // Blocks till atleast one case can run
		// Add to channel c
		case c <- x:
			x, y = y, x+y
		// Check channel quit
		case <-quit:
			fmt.Println("quit")
			return
		default:
			time.Sleep(50 * time.Millisecond) // Blocks adding to c
		}
	}
}

func routinesAndChannels() {

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go chanSum(s[:len(s)/2], c) // Run as routine
	go chanSum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	// Channel with Buffer
	buffCh := make(chan int, 2)
	buffCh <- 1
	buffCh <- 2
	fmt.Println(<-buffCh)
	fmt.Println(<-buffCh)

	// Range and Close Channel
	fiboChan := make(chan int, 10)
	go fibonacciChannel(cap(fiboChan), fiboChan)
	for i := range fiboChan {
		fmt.Println(i)
	}

	// Select across Channels
	selChan := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-selChan) // Drains 10 items from channel
		}
		quit <- 0 // Add number to quit channel to quit after 10 drains
	}()
	fibonacciSelect(selChan, quit) // Keeps adding to channel infinitely

}

// Sync Using Mutex
// SafeCounter is safe to use concurrently.
type safeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *safeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *safeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func syncUsingMuxtex() {
	c := safeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

func main() {

	// Hello World
	fmt.Println("Hello World")

	// Assign and add ints
	fmt.Println("Add nums ")
	x := 5
	y := 7
	sum := x + y
	fmt.Println(sum)

	// Data structs
	fmt.Println("Data Structs")
	dataStructs()

	// Programming constructs
	fmt.Println("Programming constructs")
	programmingConstructs()

	// Functions and Pointers
	fmt.Println("Functions and Pointers")
	functionsAndPointers()

	pic.Show(Pic)

	fmt.Println("Routines and Channels")
	routinesAndChannels()

	fmt.Println("Sync Using Mutex")
	syncUsingMuxtex()

}
