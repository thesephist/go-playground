package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("what")
		} else {
			fmt.Println("How")
		}
	}

	// we can declare extra things in preconditions
	if j := 2; 10%j == 0 {
		fmt.Println("10 is divisible by", j)
	}
}
