package main

import (
	"bufio"
	"fmt"
	"os"
)

func hc(m []string) int {
	for i := range m {
		if i+1 < len(m) && m[i] == m[i+1] {
			l, r := i, i+1
			for r < len(m) && l >= 0 && m[l] == m[r] {
				l, r = l-1, r+1
			}
			l, r = l+1, r-1
			if l == 0 || r == len(m)-1 {
				return l + (r-l)/2 + 1
			}
		}
	}
	return 0
}

func vc(m []string) int {
	nm := []string{}
	for x := range m[0] {
		s := ""
		for y := range m {
			s += string(m[y][x])
		}
		rs := ""
		for i := len(s) - 1; i >= 0; i-- {
			rs += string(s[i])
		}
		nm = append(nm, rs)
	}
	return hc(nm)
}

func main() {
	res := 0
	m := []string{}
	ms := [][]string{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			ms = append(ms, m)
			m = []string{}
		} else {
			m = append(m, l)
		}
	}
	ms = append(ms, m)
	for _, m := range ms {
		res += vc(m)
		res += 100 * hc(m)
	}
	fmt.Println(res)
}
