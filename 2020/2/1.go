package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	valid := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		minMax := strings.Split(split[0], "-")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])
		letter := string(split[1][0])
		password := split[2]
		count := strings.Count(password, letter)
		if count >= min && count <= max {
			valid++
		}
	}
	fmt.Println(valid)
}
