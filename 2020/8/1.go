package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	accumulator := 0

	var ops [][]string

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Split(line, " ")
		ops = append(ops, []string{args[0], string(args[1][0]), args[1][1:], "0"})
	}

	for i := 0; i < len(ops); i++ {
		val, _ := strconv.Atoi(ops[i][2])
		if ops[i][3] == "0" {
			ops[i][3] = "1"
			if ops[i][0] == "acc" {
				if ops[i][1] == "+" {
					accumulator += val
				} else {
					accumulator -= val
				}
			} else if ops[i][0] == "jmp" {
				if ops[i][1] == "+" {
					i += val - 1
				} else {
					i -= val + 1
				}
			}
		} else {
			break
		}
	}
	fmt.Println(accumulator)
}
