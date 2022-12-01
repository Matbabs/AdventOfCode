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
	questions := make(map[rune]bool)

	setCount := func() {
		count += len(questions)
		questions = make(map[rune]bool)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			setCount()
			continue
		}
		for _, c := range line {
			if c != ' ' {
				questions[c] = true
			}
		}

	}
	setCount()
	fmt.Println(count)
}
