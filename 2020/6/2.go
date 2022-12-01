package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	count := 0
	questions := make(map[rune]int)
	people := 0

	setCount := func() {
		for _, v := range questions {
			if v == people {
				count++
			}
		}
		questions = make(map[rune]int)
		people = 0
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			setCount()
			continue
		}
		people++
		for _, c := range line {
			if c != ' ' {
				questions[c] += 1
			}
		}
	}
	setCount()
	fmt.Println(count)
}
