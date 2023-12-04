package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
)

func main() {
	res := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("\\d+")
	for scanner.Scan() {
		nbrs := re.FindAllString(scanner.Text(), -1)
		win, yrs := nbrs[1:11], nbrs[11:]
		ct := 0
		for _, y := range yrs {
			if slices.Contains(win, y) {
				ct++
			}
		}
		if ct > 0 {
			res += 1 << (ct - 1)
		}
	}
	fmt.Println(res)
}
