package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	steps := strings.Split(string(input), "\n")
	seenFreqs := map[int]int{0: 1}
	freq := 0
	for i := 0; ; i++ {
		s := steps[i%len(steps)]
		if len(s) < 2 {
			continue
		}
		action := s[0]
		num, err := strconv.Atoi(s[1:])
		if err != nil {
			fmt.Errorf("Invalid input given %s", s)
		}

		switch action {
		case '+':
			freq += num
		case '-':
			freq -= num
		}

		seenFreqs[freq]++
		if seenFreqs[freq] > 1 {
			break
		}
	}

	fmt.Println("Problem 1 Part 2 Answer:", freq)
}
