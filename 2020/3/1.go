package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	position := 0
	count := 0

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if position < len(line) && line[position] == '#' {
			count++
		}
		position += 3
		position = position % len(line)

	}

	fmt.Println(count)
}
