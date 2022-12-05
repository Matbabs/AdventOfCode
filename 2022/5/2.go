package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var stacks = make(map[int][]string)
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
					stacks[int(i)] = append(stacks[int(i)], string(c))
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
		tmp := make([]string, len(stacks[op[1]][0:op[0]]))
		copy(tmp, stacks[op[1]][0:op[0]])
		stacks[op[2]] = append(tmp, stacks[op[2]]...)
		stacks[op[1]] = stacks[op[1]][op[0]:]
	}
	res := ""
	for i := 0; i < len(stacks); i++ {
		res += stacks[i+1][0]
	}
	fmt.Println(res)
}
