package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

const offset int = 32 // integer delta between capital/lowercase letter in ascii

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func cmp(x, y byte) bool {
	return abs(int(x)-int(y)) == offset
}

func collapse(input []byte) int {
	for x := 0; x < len(input)-1; {
		if cmp(input[x], input[x+1]) {
			input = append(input[:x], input[x+2:]...)
			x = 0
		} else {
			x++
		}
	}
	return len(input)
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = bytes.TrimRight(input, "\n")
	answer := collapse(input)
	fmt.Printf("Problem 5.1 Answer: %d\n", answer)
}
