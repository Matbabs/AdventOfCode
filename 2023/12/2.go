package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var memo map[string]int

func cbn(p string, i int, a string, c []int) int {
	cst := ""
	for cs := range c {
		cst += "," + strconv.Itoa(cs)
	}
	key := fmt.Sprintf("%s:%d:%s%s", p, i, a, cst)
	if v, ok := memo[key]; ok {
		return v
	}
	re := regexp.MustCompile("#+")
	gs := re.FindAllString(a, -1)
	if len(gs) > 1 && len(gs) <= len(c) {
		for k := range gs[:len(gs)-1] {
			if len(gs[k]) != c[k] {
				memo[key] = 0
				return 0
			}
		}
	}
	if i == len(p) {
		if len(gs) == len(c) {
			for k := range gs {
				if len(gs[k]) != c[k] {
					memo[key] = 0
					return 0
				}
			}
			if len(gs[len(gs)-1]) == c[len(c)-1] {
				memo[key] = 1
				return 1
			}
		}
		memo[key] = 0
		return 0
	}
	ct := 0
	if p[i] == '?' {
		ct += cbn(p, i+1, a+"#", c)
		ct += cbn(p, i+1, a+".", c)
	} else {
		ct += cbn(p, i+1, a+string(p[i]), c)
	}
	memo[key] = ct
	return memo[key]
}

func main() {
	exp := 5
	res := 0
	p := []string{}
	a := [][]int{}
	re := regexp.MustCompile("\\d+")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := strings.Split(scanner.Text(), " ")
		ps := ""
		for i := 0; i < exp; i++ {
			if i != exp-1 {
				ps += l[0] + "?"
			} else {
				ps += l[0]
			}
		}
		p = append(p, ps)
		n := re.FindAllString(l[1], -1)
		sa := []int{}
		for i := 0; i < exp; i++ {
			for _, ds := range n {
				d, _ := strconv.Atoi(ds)
				sa = append(sa, d)
			}
		}
		a = append(a, sa)
	}
	re = regexp.MustCompile("#+")
	for i := range p {
		memo = make(map[string]int)
		res += cbn(p[i], 0, "", a[i])
	}
	fmt.Println(res)
}
