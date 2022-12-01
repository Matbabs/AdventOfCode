package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		max_1    = 0
		max_2    = 0
		max_3    = 0
		calories = 0
	)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if calories > max_1 {
				max_3, max_2, max_1 = max_2, max_1, calories
			} else if calories > max_2 {
				max_3, max_2 = max_2, calories
			} else if calories > max_3 {
				max_3 = calories
			}
			calories = 0
			continue
		}
		kcl, _ := strconv.Atoi(line)
		calories += kcl
	}
	fmt.Println(max_1 + max_2 + max_3)
}
