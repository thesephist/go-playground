package main

import (
	"fmt"
	"os"
	"strconv"
)

type factorList []int

func (factors factorList) String() string {
	s := ""
	for _, f := range factors {
		s += strconv.Itoa(f) + " * "
	}
	return s[:len(s)-3]
}

func primeFactorize(numberToFactorize int) (factors factorList) {
	n := numberToFactorize
	k := 2 // first prime to attempt
	factors = []int{}
	for k < numberToFactorize {
		for n%k == 0 {
			n /= k
			factors = append(factors, k)
			if n == 1 {
				return
			}
		}
		k++
	}
	// should never happen
	return
}

func main() {
	input := os.Args[1]
	numberToFactorize, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println(input, "isn't a valid number to factorize!")
		os.Exit(1)
	}

	factors := primeFactorize(numberToFactorize)
	fmt.Println("Factorized", numberToFactorize, "->", factors)
}
