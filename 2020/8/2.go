package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func program(ops [][]string, fix int) (int, bool) {
	accumulator := 0
	ops_copy := make([][]string, len(ops))
	for i := range ops {
		ops_copy[i] = make([]string, len(ops[i]))
		copy(ops_copy[i], ops[i])
	}
	if ops_copy[fix][0] == "jmp" {
		ops_copy[fix][0] = "nop"
	} else {
		ops_copy[fix][0] = "jmp"
	}
	for i := 0; i < len(ops_copy); i++ {
		val, _ := strconv.Atoi(ops_copy[i][2])
		if ops_copy[i][3] == "0" {
			ops_copy[i][3] = "1"
			if ops_copy[i][0] == "acc" {
				if ops_copy[i][1] == "+" {
					accumulator += val
				} else {
					accumulator -= val
				}
			} else if ops_copy[i][0] == "jmp" {
				if ops_copy[i][1] == "+" {
					i += val - 1
				} else {
					i -= val + 1
				}
			}
		} else {
			return accumulator, true
		}
	}
	return accumulator, false
}

func main() {
	var ops [][]string
	var jmps_nops []int

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Split(line, " ")
		ops = append(ops, []string{args[0], string(args[1][0]), args[1][1:], "0"})
		if args[0] == "jmp" || args[0] == "nop" {
			jmps_nops = append(jmps_nops, len(ops)-1)
		}
	}
	for i := 0; i < len(jmps_nops); i++ {
		res, err := program(ops, jmps_nops[i])
		if err == false {
			fmt.Println(res)
			break
		}
	}
}
