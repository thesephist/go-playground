package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		key, val := pair[0], pair[1]

		switch key {
		case "PATH":
			fmt.Printf("%s ->", key)
			paths := strings.Split(val, ":")
			for _, path := range paths {
				fmt.Println("\t" + path)
			}
		default:
			fmt.Printf("%s -> %s\n", key, val)
		}

	}
}
