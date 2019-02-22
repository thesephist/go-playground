package main

import "fmt"

// constants can't be declared a type with := becuase
//  := syntax is for variables. Think of const declarations
//  as #define in C.
const s1, s2 string = "Hello, ", "Hi"

func main() {
	fmt.Println("Hello")
	fmt.Println(s1 + s2)

	// numeric constants aren't given a type until first use
	const num = 1000.0
	fmt.Println("num", num/3)
}
