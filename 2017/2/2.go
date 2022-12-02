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
		for i := 0; i < len(numbers); i++ {
			for j := i + 1; j < len(numbers); j++ {
				a, _ := strconv.Atoi(numbers[i])
				b, _ := strconv.Atoi(numbers[j])
				if b > a {
					a, b = b, a
				}
				if a%b == 0 {
					checksum += a / b
				}
			}
		}
	}
	fmt.Println(checksum)
}
