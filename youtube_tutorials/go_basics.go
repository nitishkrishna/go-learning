package main

import (
	"errors"
	"fmt"
	"math"
	"runtime"
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

func main() {

	// Hello World
	fmt.Println("Hello World")

	// Assign and add ints
	fmt.Println("Add nums ")
	x := 5
	y := 7
	sum := x + y
	fmt.Println(sum)

	// If condition
	fmt.Println("If cond ")
	if x > 6 {
		fmt.Println("More than 6")
	} else if x == 5 {
		fmt.Println("Equal to 5")
	} else {
		fmt.Println("Less than 6, Not 5")
	}

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

	// Creating struct of a type
	fmt.Println("Structs")
	p := person{name: "Nitish", age: 45}
	fmt.Println(p)
	fmt.Println("Age:", p.age)
	// Pointers to structs
	pPtr := &p
	personPtr := &person{"Alice", 55}
	// Doesn't need *
	pPtr.age = 50
	fmt.Println(p, personPtr)

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

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

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

}
