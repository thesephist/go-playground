package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}
	str := "Hello, World!"

	// iterate over array / slice
	for idx, val := range slc {
		fmt.Println(val, "value at index", idx)
	}

	// iterate over key values in a map
	for key, val := range m {
		fmt.Println(key, "->", val)
	}

	// iterate over Unicode code points in a string
	for idx, char := range str {
		fmt.Println("character", string(char), "at idx", idx)
	}
}
