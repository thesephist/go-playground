package main

import "fmt"
import "strconv"

func main() {
	fmt.Println("Hi")

	var strings []string
	for len(strings) < 30 {
		count := strconv.Itoa(len(strings))
		strings = append(strings, "x"+count)
	}
	fmt.Println(strings)

	// make a zeroed-out slice
	newStrings := make([]string, len(strings))
	copy(newStrings, strings)
	fmt.Println("New:", newStrings)

	// slicing slices
	fmt.Println("slice:", strings[10:20])

	for _, s := range strings {
		fmt.Println("includes", s)
	}
}
