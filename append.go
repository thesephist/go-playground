package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// read which file to read
	fileNamePtr := flag.String("file", "", "file to echo")
	newLinePtr := flag.String("line", "", "line to append to file")

	flag.Parse()

	if *fileNamePtr == "" {
		fmt.Println("Please pass a file to read")
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(*fileNamePtr)

	fileContents := []byte(string(data) + *newLinePtr + "\n")
	err = ioutil.WriteFile(*fileNamePtr, fileContents, 0644)
	check(err)
}
