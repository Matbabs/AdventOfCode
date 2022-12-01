package main

import (
	"bufio"
	"fmt"
	"os"
)

func slope(right int, down int) int {
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
		position += right
		if down > 1 {
			for i := 1; i < down; i++ {
				scanner.Scan()
			}
		}
		position = position % len(line)
	}
	return count
}

func main() {
	fmt.Println(slope(1, 1) * slope(3, 1) * slope(5, 1) * slope(7, 1) * slope(1, 2))
}
