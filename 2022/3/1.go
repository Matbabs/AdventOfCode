package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sum := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bag := scanner.Text()
		left := bag[:len(bag)/2]
		right := bag[len(bag)/2:]
		for i := 0; i < len(left); i++ {
			if strings.Count(right, string(left[i])) > 0 {
				if left[i] >= 'a' {
					sum += int(left[i] - 'a' + 1)
				} else {
					sum += int(left[i] - 'A' + 27)
				}
				break
			}
		}
	}
	fmt.Println(sum)
}
