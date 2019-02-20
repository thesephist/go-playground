package main

import "fmt"

type walker interface {
	walk()
}

type human struct {
	walked bool
}

type robot struct {
	walked bool
}

func (h *human) walk() {
	h.walked = true
	fmt.Println("Human is walking. Step step step.")
}

func (r *robot) walk() {
	r.walked = true
	fmt.Println("Robot is walking. Bleep bloop bleep.")
}

func doWalk(w walker) {
	w.walk()
	fmt.Println("Walked!")
}

func main() {
	h := *new(human)
	r := *new(robot)

	fmt.Println(h, r)

	// here, we're passing the pointer to a struct
	//  into the interface; the value is still mutable
	//  and the struct is automaticlaly dereferenced, but
	//  doWalk() takes a walker, not a *walker.
	// Go interfaces can hold either a struct or a pointer to a struct,
	//  and they behave differently.
	doWalk(&h)
	doWalk(&r)

	fmt.Println(h, r)
}
