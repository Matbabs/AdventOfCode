package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isPlayable(hand string, colors map[string]int, re *regexp.Regexp) bool {
	for _, h := range strings.Split(hand, ";") {
		nbs := make(map[string]int)
		for _, s := range strings.Split(h, ",") {
			n, _ := strconv.Atoi(re.FindString(s))
			for c := range colors {
				if strings.Contains(s, c) {
					nbs[c] += n
					if nbs[c] > colors[c] {
						return false
					}
				}
			}
		}
	}
	return true
}

func main() {
	res := 0
	colors := map[string]int{"red": 12, "green": 13, "blue": 14}
	re := regexp.MustCompile("\\d+")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		g := strings.Split(scanner.Text(), ":")
		gn, _ := strconv.Atoi(re.FindString(g[0]))
		if isPlayable(g[1], colors, re) {
			res += gn
		}
	}
	fmt.Println(res)
}
