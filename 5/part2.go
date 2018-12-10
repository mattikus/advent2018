package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
)

const offset byte = 32 // integer delta between capital/lowercase letter in ascii

func cmp(a, b byte) bool {
	return a != b && (a|offset == b || b|offset == a)
}

func collapse(input []byte) int {
	for i := 0; i < len(input)-1; i++ {
		if cmp(input[i], input[i+1]) {
			input = append(input[:i], input[i+2:]...)
			i = -1
		}
	}
	return len(input)
}

func collapseWorker(input []byte, jobs <-chan byte, results chan<- int) {
	for j := range jobs {
		polymer := bytes.Replace(input, []byte{j}, []byte{}, -1)
		polymer = bytes.Replace(polymer, []byte{j + offset}, []byte{}, -1)
		results <- collapse(polymer)
	}
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = bytes.TrimRight(input, "\n")

	results := make(chan int, 26)
	jobs := make(chan byte, 26)

	// Build a worker pool roughly equal to the cores available
	for w := 0; w < runtime.NumCPU(); w++ {
		go collapseWorker(input, jobs, results)
	}

	// Create some jobs where we remove each unit from the polymer
	for j := byte('A'); j < byte('a'); j++ {
		jobs <- j
	}
	close(jobs)

	var answer int
	for i := 0; i < 26; i++ {
		if r := <-results; r < answer || answer == 0 {
			answer = r
		}
	}

	fmt.Printf("Problem 5.2 Answer: %d\n", answer)
}
