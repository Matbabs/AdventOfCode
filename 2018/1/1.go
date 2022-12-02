package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	frequency := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		op := scanner.Text()
		v, _ := strconv.Atoi(op[1:])
		if op[0] == '+' {
			frequency += v
		} else {
			frequency -= v
		}
	}
	fmt.Println(frequency)
}
