package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	res := 0
	colors := map[string]int{"red": 12, "green": 13, "blue": 14}
	re := regexp.MustCompile("\\d+")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		g := strings.Split(scanner.Text(), ":")
		maxNbs := make(map[string]int)
		for _, h := range strings.Split(g[1], ";") {
			nbs := make(map[string]int)
			for _, s := range strings.Split(h, ",") {
				n, _ := strconv.Atoi(re.FindString(s))
				for c := range colors {
					if strings.Contains(s, c) {
						nbs[c] += n
						if nbs[c] > maxNbs[c] {
							maxNbs[c] = nbs[c]
						}
					}
				}
			}
		}
		subRes := 1
		for _, v := range maxNbs {
			subRes *= v
		}
		res += subRes
	}
	fmt.Println(res)
}
