package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var stacks = make(map[int][]rune)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		i := float32(0.5)
		for j, c := range line {
			if (j-1)%2 == 0 {
				i += 0.5
				if c >= 65 && c <= 90 {
					stacks[int(i)] = append(stacks[int(i)], c)
				}
			}
		}
	}
	re := regexp.MustCompile(`\d+`)
	for scanner.Scan() {
		op := make([]int, 3)
		for i, n := range re.FindAllString(scanner.Text(), 3) {
			op[i], _ = strconv.Atoi(n)
		}
		for i := 0; i < op[0]; i++ {
			stacks[op[2]] = append([]rune{stacks[op[1]][0]}, stacks[op[2]]...)
			stacks[op[1]] = stacks[op[1]][1:]
		}
	}
	res := ""
	for i := 0; i < len(stacks); i++ {
		res += string(stacks[i+1][0])
	}
	fmt.Println(res)
}
