package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func program(codes []int, noun int, verb int) int {
	codes_copy := make([]int, len(codes))
	copy(codes_copy, codes)
	codes_copy[1] = noun
	codes_copy[2] = verb
	for i := 0; i < len(codes_copy); i += 4 {
		switch codes_copy[i] {
		case 1:
			codes_copy[codes_copy[i+3]] = codes_copy[codes_copy[i+1]] + codes_copy[codes_copy[i+2]]
		case 2:
			codes_copy[codes_copy[i+3]] = codes_copy[codes_copy[i+1]] * codes_copy[codes_copy[i+2]]
		case 99:
			break
		}
	}
	return codes_copy[0]
}

func main() {

	var codes []int

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		for _, n_str := range strings.Split(scanner.Text(), ",") {
			n_int, _ := strconv.Atoi(n_str)
			codes = append(codes, n_int)
		}
	}

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			if program(codes, noun, verb) == 19690720 {
				fmt.Println(100*noun + verb)
				return
			}
		}
	}
}
