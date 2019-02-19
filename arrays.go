package main

import "fmt"

func main() {
	// empty array
	var a [5]int
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// valued array
	var b = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		fmt.Println(b[i])
	}

	fmt.Println("Printing whole arrays:")
	fmt.Println(a, b)
}
