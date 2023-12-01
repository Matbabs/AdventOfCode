package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	res := 0
	wToNb := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var nbs []string
		line := scanner.Text()
		for i, c := range line {
			if unicode.IsDigit(c) {
				nbs = append(nbs, string(c))
			}
			for w, n := range wToNb {
				if i+len(w) <= len(line) && line[i:i+len(w)] == w {
					nbs = append(nbs, n)
				}
			}
		}
		ns, _ := strconv.Atoi(nbs[0] + nbs[len(nbs)-1])
		res += ns
	}
	fmt.Println(res)
}
