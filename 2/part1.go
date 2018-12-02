package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	twos := 0
	threes := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		counts := map[rune]int{}
		for _, i := range scanner.Text() {
			counts[i]++
		}

		vals := map[int]struct{}{}
		for _, c := range counts {
			vals[c] = struct{}{}
		}

		if _, ok := vals[2]; ok {
			twos++
		}

		if _, ok := vals[3]; ok {
			threes++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("error reading input:", err)
	}

	fmt.Println("Problem 2 Part 1 Answer:", twos*threes)
}
