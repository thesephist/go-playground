package main

import "fmt"

func add(a, b, c int) int {
	return a + b + c
}

func sumAndDifference(a, b int) (int, int) {
	return a + b, a - b
}

func sumVariadic(addends ...int) int {
	sum := 0
	for _, val := range addends {
		sum += val
	}
	return sum
}

func main() {
	fmt.Println("Sum is", add(9, 1, 2))

	fmt.Println("With 10 and 7:")
	sum, diff := sumAndDifference(10, 7)
	fmt.Println("Sum is", sum)
	fmt.Println("Difference is", diff)

	fmt.Println("Sum of a bunch of numbers")
	fmt.Println("Sum to 10", sumVariadic(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	// spread operator
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum to 5", sumVariadic(numbers...))
}
