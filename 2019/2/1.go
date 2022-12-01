package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	codes[1] = 12
	codes[2] = 2

	for i := 0; i < len(codes); i += 4 {
		switch codes[i] {
		case 1:
			codes[codes[i+3]] = codes[codes[i+1]] + codes[codes[i+2]]
		case 2:
			codes[codes[i+3]] = codes[codes[i+1]] * codes[codes[i+2]]
		case 99:
			break
		}
	}
	fmt.Println(codes[0])
}
