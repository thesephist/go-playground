package main

import (
	"fmt"
	"time"
)

func say(msg string, sink chan<- string) {
	time.Sleep(1 * time.Second)
	sink <- msg
}

func main() {
	sink1 := make(chan string)
	sink2 := make(chan string)

	go say("Hello", sink1)
	go say("Goodbye!", sink2)

	// non-deterministic thread scheduling!
	select {
	case msg := <-sink1:
		fmt.Println("From sink 1:", msg)
	case msg := <-sink2:
		fmt.Println("From sink 2:", msg)
	}
}
