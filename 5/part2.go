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

func removeCollapse(input []byte, cut byte) int {
	input = bytes.Replace(input, []byte{cut}, []byte{}, -1)
	input = bytes.Replace(input, []byte{byte(int(cut) + offset)}, []byte{}, -1)
	return collapse(input)
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input = bytes.TrimRight(input, "\n")
	answer := 100000000
	for x := 65; x < 91; x++ {
		if num := removeCollapse(input, byte(x)); num < answer {
			answer = num
		}
	}

	fmt.Printf("Problem 5.2 Answer: %d\n", answer)
}
