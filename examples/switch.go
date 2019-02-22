package main

import "fmt"
import "time"

func main() {
	fmt.Println("Hello")

	now := time.Now()
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())

	for i := 0; i < 100; i++ {
		switch i {
		case 10, 20, 30, 40, 50, 60, 70:
			fmt.Println("Caught it!", i)
		default:
			fmt.Println("Missed it", i)
		}
	}
}
