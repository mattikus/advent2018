package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
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

func collapseWorker(jobs <-chan []byte, results chan<- int) {
	for j := range jobs {
		results <- collapse(j)
	}
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input = bytes.TrimRight(input, "\n")
	results := make(chan int, 26)
	jobs := make(chan []byte, 26)

	// Build a worker pool roughly equal to the cores available
	for w := 1; w <= runtime.NumCPU(); w++ {
		go collapseWorker(jobs, results)
	}

	for j := 65; j < 91; j++ {
		jobs <- bytes.Replace(input, []byte{byte(j)}, []byte{}, -1)
	}
	close(jobs)

	answer := 100000000
	for i := 0; i <= 25; i++ {
		answer = min(answer, <-results)
	}

	fmt.Printf("Problem 5.2 Answer: %d\n", answer)
}
