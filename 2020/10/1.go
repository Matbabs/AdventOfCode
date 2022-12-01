package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	var adapters []int
	count_1 := 1
	count_3 := 1

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		adapters = append(adapters, val)
	}

	sort.Ints(adapters)

	for i := 0; i < len(adapters)-1; i++ {
		if adapters[i+1]-adapters[i] == 1 {
			count_1++
		} else if adapters[i+1]-adapters[i] == 3 {
			count_3++
		}
	}

	fmt.Println(count_1 * count_3)
}
