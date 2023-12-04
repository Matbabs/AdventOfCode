package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	res := 0
	cds := make(map[int]int)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("\\d+")
	for scanner.Scan() {
		nbrs := re.FindAllString(scanner.Text(), -1)
		c, _ := strconv.Atoi(nbrs[0])
		win, yrs := nbrs[1:11], nbrs[11:]
		cds[c]++
		ct := 0
		for _, y := range yrs {
			if slices.Contains(win, y) {
				ct++
			}
		}
		for cp := 0; cp < cds[c]; cp++ {
			for i := 1; i <= ct; i++ {
				cds[c+i]++
			}
		}
	}
	for _, c := range cds {
		res += c
	}
	fmt.Println(res)
}
