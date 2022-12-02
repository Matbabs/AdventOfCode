package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	two := 0
	three := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count := func(s string, n int) bool {
		for _, c := range s {
			if strings.Count(s, string(c)) == n {
				return true
			}
		}
		return false
	}
	for scanner.Scan() {
		id := scanner.Text()
		if count(id, 2) {
			two++
		}
		if count(id, 3) {
			three++
		}
	}
	println(two * three)
}
