package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sum := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		digits := scanner.Text()
		for i := range digits {
			if digits[i] == digits[(i+len(digits)/2)%len(digits)] {
				sum += int(digits[i] - '0')
			}
		}
	}
	fmt.Println(sum)
}
