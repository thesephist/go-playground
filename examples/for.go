package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("%d -> ", i)
	}
	fmt.Println("10")

	j := "a"
	for len(j) < 8 {
		fmt.Println(j)
		j += "a"
	}
}
