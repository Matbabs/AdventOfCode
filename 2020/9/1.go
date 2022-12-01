package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	n := 0
	preamble := 25
	var numbers []int

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, num)
		if n > preamble {
			pairs := make(map[int]bool)
			for i := n - preamble; i < n; i++ {
				for j := i + 1; j < n; j++ {
					pairs[numbers[i]+numbers[j]] = true
				}
			}
			if !pairs[numbers[n]] {
				fmt.Println(numbers[n])
				break
			}
		}
		n++
	}
}
