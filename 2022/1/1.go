package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	max_calories := 0
	calories := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if calories > max_calories {
				max_calories = calories
			}
			calories = 0
			continue
		}
		kcl, _ := strconv.Atoi(line)
		calories += kcl
	}
	fmt.Println(max_calories)
}
