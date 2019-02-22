package main

import "fmt"

func run(count int, str string) {
	for i := 0; i < count; i++ {
	}
	fmt.Println(str)
}

func main() {
	go run(100000, "Hello")
	go run(900, "Good day!")
	go run(10000, "Goodbye!")

	// main exits when main procedure calls are done, even if not all
	//  goroutines finish. This Scanln() call waits for user input before
	//  main routine exits.
	fmt.Scanln()
	fmt.Println("done")
}
