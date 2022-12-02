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
		for i, c := range scanner.Text() {
			if c == '(' {
				floor++
			} else {
				floor--
			}
			if floor == -1 {
				fmt.Println(i + 1)
				return
			}
		}
	}
}
