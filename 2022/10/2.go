package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	X := 1
	cycle := 1
	offset := 40
	stack := make([]int, 0)
	screen := " "
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		op := strings.Split(line, " ")
		stack = append(stack, 0)
		if op[0] == "addx" {
			v, _ := strconv.Atoi(op[1])
			stack = append(stack, v)
		}
	}
	for len(stack) > 0 {
		X += stack[0]
		stack = stack[1:]
		if cycle%offset >= X-1 && cycle%offset <= X+1 {
			screen += "#"
		} else {
			screen += " "
		}
		fmt.Print(string(screen[cycle-1]))
		if cycle%offset == 0 {
			fmt.Println()
		}
		cycle++
	}
}
