package main

import "fmt"
import "time"

func report(msg string, host chan<- string) {
	time.Sleep(2 * time.Second)
	host <- msg
}

func main() {
	// buffered to 3 messages
	hostChannel := make(chan string, 3)

	go report("Hello", hostChannel)
	fmt.Println(<-hostChannel)

	go report("Work it!", hostChannel)
	go report("Goodbye", hostChannel)
	fmt.Println(<-hostChannel)
	fmt.Println(<-hostChannel)
}
