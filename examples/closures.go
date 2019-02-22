package main

import "fmt"

func newSource() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println("Printing from tow difference sources")

	source1 := newSource()
	source2 := newSource()

	fmt.Println(source1())
	fmt.Println(source2())

	fmt.Println(source1())
	fmt.Println(source2())

	fmt.Println(source1())
	fmt.Println(source2())
}
