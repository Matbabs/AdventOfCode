package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func cbn(p string, i int, a string, c []int, cb *[]string) {
	re := regexp.MustCompile("#+")
	gs := re.FindAllString(a, -1)
	if len(gs) > 1 && len(gs) <= len(c) {
		for k := range gs[:len(gs)-1] {
			if len(gs[k]) != c[k] {
				return
			}
		}
	}
	if i == len(p) {
		if len(gs) == len(c) {
			for k := range gs {
				if len(gs[k]) != c[k] {
					return
				}
			}
			if len(gs[len(gs)-1]) == c[len(c)-1] {
				*cb = append(*cb, a)
				return
			}
		}
		return
	}
	if p[i] == '?' {
		cbn(p, i+1, a+"#", c, cb)
		cbn(p, i+1, a+".", c, cb)
	} else {
		cbn(p, i+1, a+string(p[i]), c, cb)
	}
}

func main() {
	res := 0
	p := []string{}
	a := [][]int{}
	re := regexp.MustCompile("\\d+")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := strings.Split(scanner.Text(), " ")
		p = append(p, l[0])
		n := re.FindAllString(l[1], -1)
		sa := []int{}
		for _, ds := range n {
			d, _ := strconv.Atoi(ds)
			sa = append(sa, d)
		}
		a = append(a, sa)
	}
	re = regexp.MustCompile("#+")
	for i := range p {
		var cb []string
		cbn(p[i], 0, "", a[i], &cb)
		res += len(cb)
	}
	fmt.Println(res)
}
