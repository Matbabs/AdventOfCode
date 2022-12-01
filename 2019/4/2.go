package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPassword(password int) bool {
	var digits []int
	adjacent := false
	adjs := make(map[int]int)
	for password > 0 {
		digits = append(digits, password%10)
		password /= 10
	}
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] < digits[i+1] {
			return false
		}
		if digits[i] == digits[i+1] {
			adjacent = true
			adjs[digits[i]]++
		}
	}
	if !adjacent {
		return false
	}
	for _, v := range adjs {
		if v == 1 {
			return true
		}
	}
	return false
}

func main() {
	count := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bounds := strings.Split(line, "-")
		min, _ := strconv.Atoi(bounds[0])
		max, _ := strconv.Atoi(bounds[1])
		for i := min; i <= max; i++ {
			if isPassword(i) {
				count++
			}
		}
	}
	fmt.Println(count)
}
