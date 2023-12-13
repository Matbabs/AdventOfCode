package main

import (
	"bufio"
	"fmt"
	"os"
)

func hc(m []string, f ...int) (int, int) {
	for i := range m {
		if i+1 < len(m) && m[i] == m[i+1] {
			l, r := i, i+1
			for r < len(m) && l >= 0 && m[l] == m[r] {
				l, r = l-1, r+1
			}
			l, r = l+1, r-1
			if (l == 0 || r == len(m)-1) && (len(f) == 0 || f[0] != i) {
				return l + (r-l)/2 + 1, i
			}
		}
	}
	return 0, 0
}

func vc(m []string, f ...int) (int, int) {
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
	return hc(nm, f...)
}

func cb(m []string) (int, int) {
	ov, vi := vc(m)
	oh, hi := hc(m)
	for i := range m {
		for j := range m[0] {
			nm := []string{}
			for k := range m {
				s := ""
				for l := range m[0] {
					if i == k && j == l {
						if m[i][j] == '.' {
							s += "#"
						} else {
							s += "."
						}
					} else {
						s += string(m[k][l])
					}
				}
				nm = append(nm, s)
			}
			v, _ := vc(nm)
			h, _ := hc(nm)
			if v == ov && h == oh {
				v, _ = vc(nm, vi)
				h, _ = hc(nm, hi)
			}
			if (ov != v || oh != h) && !(v == 0 && h == 0) {
				if ov != v {
					return v, 0
				} else {
					return 0, h
				}
			}
		}
	}
	return 0, 0
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
		v, h := cb(m)
		res += v
		res += 100 * h
	}
	fmt.Println(res)
}
