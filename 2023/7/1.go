package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func hand(c string) string {
	t := make(map[rune]int)
	for _, r := range c {
		t[r]++
	}
	switch len(t) {
	case 1:
		return "fv"
	case 2:
		for _, v := range t {
			if v == 4 {
				return "fr"
			}
		}
		if len(c) != 3 {
			return "fl"
		}
	case 3:
		for _, v := range t {
			if v == 3 {
				return "t"
			}
		}
		if len(c) != 3 {
			return "ps"
		}
	case 4:
		return "p"
	}
	return "o"
}

func main() {
	res := 0
	tp := map[string]int{"fv": 6, "fr": 5, "fl": 4, "t": 3, "ps": 2, "p": 1, "o": 0}
	cd := map[string]int{"2": 0, "3": 1, "4": 2, "5": 3, "6": 4, "7": 5, "8": 6, "9": 7, "T": 8, "J": 9, "Q": 10, "K": 11, "A": 12}
	cds := make([][]string, 0)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := strings.Split(scanner.Text(), " ")
		c = append(c, hand(c[0]))
		cds = append(cds, c)
	}
	sort.SliceStable(cds, func(i, j int) bool {
		if cds[i][2] != cds[j][2] {
			return tp[cds[i][2]] < tp[cds[j][2]]
		}
		d := 0
		for k := 0; k < len(cds[i][0]); k++ {
			if cds[i][0][k] != cds[j][0][k] {
				d = k
				break
			}
		}
		return cd[string(cds[i][0][d])] < cd[string(cds[j][0][d])]
	})
	for i, c := range cds {
		b, _ := strconv.Atoi(c[1])
		res += b * (i + 1)
	}
	fmt.Println(res)
}
