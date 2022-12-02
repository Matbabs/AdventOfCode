package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	floor := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, c := range scanner.Text() {
			if c == '(' {
				floor++
			} else {
				floor--
			}
		}
	}
	fmt.Println(floor)
}
