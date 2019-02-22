package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const ANSI_BLUE string = "\u001b[34m"
const ANSI_END string = "\033[0m"

func makeExp(exp int) string {
	EXP_CHARS := []string{"⁰", "¹", "²", "³", "⁴", "⁵", "⁶", "⁷", "⁸", "⁹"}

	s := ""
	for exp > 0 {
		s, exp = EXP_CHARS[exp%10]+s, exp/10
	}
	return s
}

type factorList []int

func (factors factorList) String() string {
	parts := []string{}
	last, lastExp := 1, 1
	for _, f := range factors {
		if last == f {
			lastExp++
		} else {
			if lastExp == 1 {
				parts = append(parts, strconv.Itoa(last))
			} else {
				parts = append(parts, strconv.Itoa(last)+makeExp(lastExp))
			}
			last = f
			lastExp = 1
		}
	}
	if lastExp == 1 {
		parts = append(parts, strconv.Itoa(last))
	} else {
		parts = append(parts, strconv.Itoa(last)+makeExp(lastExp))
	}
	// 0th element is always 1 to power of 1, which is meaningless
	return strings.Join(parts[1:], " × ")
}

func primeFactorize(numberToFactorize int) (factors factorList) {
	if numberToFactorize <= 2 {
		return []int{numberToFactorize}
	}

	// potential optimization: only check when k is prime
	n, k := numberToFactorize, 2
	factors = []int{}
	for k <= numberToFactorize {
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
	inputs := os.Args[1:]
	for _, input := range inputs {
		numberToFactorize, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println(input, "isn't a valid number to factorize!")
			continue
		}

		factors := primeFactorize(numberToFactorize)
		if len(factors) == 1 {
			fmt.Printf("%d%s\tis prime%s\n", numberToFactorize, ANSI_BLUE, ANSI_END)
		} else {
			fmt.Printf("%d\t%s\n", numberToFactorize, factors)
		}
	}
}
