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
	signal := 0
	signals := 20
	offset := 40
	stack := make([]int, 0)
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
		cycle++
		X += stack[0]
		stack = stack[1:]
		if cycle == signals {
			signals += offset
			signal += X * cycle
		}
	}
	fmt.Println(signal)
}
