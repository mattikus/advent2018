package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// diffChars returns the number of characters which are different between two strings
func diffChars(a, b string) int {
	count := 0
	if len(a) != len(b) {
		return -1
	}
	for n := range a {
		if a[n] != b[n] {
			count++
		}
	}
	return count
}

// diffPos returns the first position where two strings differ, or -1 if it cannot find it
func diffPos(a, b string) int {
	if len(a) != len(b) {
		return -1
	}
	for n := 0; n < len(a); n++ {
		if a[n] != b[n] {
			return n
		}
	}
	return -1
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// get an array of boxids
	ids := []string{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("error reading input:", err)
	}

	// Iterate over them until we find the two that differ the least
	delta := 1000
	var ida, idb string
	for pos, a := range ids {
		for _, b := range ids[pos+1:] {
			d := diffChars(a, b)
			if d < 0 {
				continue
			}
			if d < delta {
				ida = a
				idb = b
				delta = d
			}
		}
	}

	// Delete characters until we're the same string
	for pos := diffPos(ida, idb); pos != -1; pos = diffPos(ida, idb) {
		ida = ida[:pos] + ida[pos+1:]
		idb = idb[:pos] + idb[pos+1:]
	}

	fmt.Printf("Problem 2 Part 1 Answer: %s\n", ida)
}
