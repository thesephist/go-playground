package main

import "fmt"

func main() {
	fmt.Println("maps!")

	dict := make(map[string]int)

	dict["first"] = 1
	dict["second"] = 2
	dict["third"] = 3

	fmt.Println(dict)

	for k, val := range dict {
		fmt.Printf("%s -> %d\n", k, val)
	}

	// make a map from scratch
	m := map[string]string{
		"one":   "first",
		"two":   "second",
		"three": "third",
		"four":  "fourth",
		"five":  "fifth",
		"six":   "sixth",
	}
	for k, val := range m {
		fmt.Printf("%s in ordinal is -> %s\n", k, val)
	}

	// detect if present
	val, present := m["six"]
	fmt.Println(val, "is present?", present)

	// delete key
	delete(m, "six")

	val, present = m["six"]
	fmt.Println(val, "is present?", present)
}
