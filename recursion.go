package main

import "fmt"

func fib(n int, memo map[int]int) int {
	// memoized fibonacci, much faster
	if n <= 1 {
		return 1
	} else {
		val, present := memo[n]
		if present {
			return val
		} else {
			memo[n] = fib(n-1, memo) + fib(n-2, memo)
			return memo[n]
		}
	}
}

func main() {
	fmt.Println("Fibonacci sequence")

	memo := make(map[int]int)
	for i := 0; i < 50; i++ {
		fmt.Printf("%d:\t%d\n", i, fib(i, memo))
	}

	fmt.Println(memo)
}
