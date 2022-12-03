package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sum := 0
	group := 1
	bags := make([]string, 3)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bags[group%3] = scanner.Text()
		if group%3 == 0 {
			for i := 0; i < len(bags[0]); i++ {
				if strings.Count(bags[1], string(bags[0][i])) > 0 && strings.Count(bags[2], string(bags[0][i])) > 0 {
					if bags[0][i] >= 'a' {
						sum += int(bags[0][i] - 'a' + 1)
					} else {
						sum += int(bags[0][i] - 'A' + 27)
					}
					break
				}
			}
		}
		group++
	}
	fmt.Println(sum)
}
