package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	freq := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		action := scanner.Text()[0]
		num, err := strconv.Atoi(scanner.Text()[1:])
		if err != nil {
			fmt.Errorf("Invalid input given %s", scanner.Text())
		}

		switch action {
		case '+':
			freq += num
		case '-':
			freq -= num
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal("error reading input:", err)
	}

	fmt.Println("Problem 1 Answer:", freq)
}
