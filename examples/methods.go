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

func (t *task) complete() {
	t.state = TASK_DONE
}

func (t *task) uncomplete() {
	t.state = TASK_INCOMPLETE
}

func (t *task) setDescription(desc string) {
	t.description = desc
}

func main() {
	t := task{
		description: "Learn how to golang good",
	}

	fmt.Println(t)

	t.complete()
	t.setDescription("Learn how to golang better!")

	fmt.Println(t)
}
