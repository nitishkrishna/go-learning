package main

import (
	"fmt"
	"sort"
)

func sortingUtil() {

	// Sort Strings in alpha order
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// Sort ints
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// Check sorted order
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted Ints: ", s)
	x := sort.StringsAreSorted(strs)
	fmt.Println("Sorted Strings: ", x)
}

func customSorting() {

}

// Create type for custom sort based on string
type byLength []string

// Override the string sort Interface funcs
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func customSort() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

func main() {
	sortingUtil()
	customSort()
}
