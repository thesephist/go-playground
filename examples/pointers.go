package main

import "fmt"

func bumpValue(i int) int {
	i++
	return i
}

func bumpPointer(i *int) int {
	*i++
	return *i
}

func main() {
	fmt.Println("Pointers!")

	numValue := 0
	numPointer := 0

	bumpValue(numValue)
	bumpPointer(&numPointer)

	fmt.Println(numValue, numPointer)
}
