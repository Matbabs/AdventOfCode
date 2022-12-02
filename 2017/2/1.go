package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	checksum := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "\t")
		v, _ := strconv.Atoi(numbers[0])
		min, max := v, v
		for _, n := range numbers[1:] {
			v, _ := strconv.Atoi(n)
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
		checksum += max - min
	}
	fmt.Println(checksum)
}
