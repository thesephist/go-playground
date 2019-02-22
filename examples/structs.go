package main

import "fmt"

type taskState int

const (
	TASK_INCOMPLETE taskState = iota
	TASK_INPROG     taskState = iota
	TASK_DONE       taskState = iota
)

type task struct {
	state       taskState
	description string
}

func main() {
	newTask := task{
		state:       TASK_INCOMPLETE,
		description: "Learn to write Go programs",
	}

	fmt.Println(newTask)
	fmt.Println(&newTask)

	fmt.Println("description is:", newTask.description)
	fmt.Println("description is (from pointer):", (&newTask).description)

	newTask.description = "Learn to write go well"

	fmt.Println("description is:", newTask.description)
}
