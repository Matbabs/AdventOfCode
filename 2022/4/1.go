package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	count := 0
	s := make([]int, 4)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		re := regexp.MustCompile(`\d+`)
		for i, n := range re.FindAllString(scanner.Text(), 4) {
			s[i], _ = strconv.Atoi(n)
		}
		if s[0] <= s[2] && s[1] >= s[3] || s[2] <= s[0] && s[3] >= s[1] {
			count++
		}
	}
	fmt.Println(count)
}
