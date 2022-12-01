package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var possibilities int64 = 0
var adapters = make(map[int]bool)
var max = 0

func visit(x int) {
	if x == max {
		possibilities++
		return
	}
	for i := 1; i <= 3; i++ {
		if _, ok := adapters[x+i]; ok {
			visit(x + i)
		}
	}
}

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		adapters[val] = true
		if val > max {
			max = val
		}
	}

	visit(0)

	fmt.Println(possibilities)
}
