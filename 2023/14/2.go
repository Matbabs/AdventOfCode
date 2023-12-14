package main

import (
	"bufio"
	"fmt"
	"os"
)

func plt(m *[][]rune) {
	for y := range *m {
		for x := range (*m)[y] {
			if (*m)[y][x] == 'O' {
				k := 1
				for y-k >= 0 && (*m)[y-k][x] == '.' {
					k++
				}
				k--
				(*m)[y][x] = '.'
				(*m)[y-k][x] = 'O'
			}
		}
	}
}

func spn(m *[][]rune) {
	nm := [][]rune{}
	for x := range (*m)[0] {
		s := []rune{}
		for y := len(*m) - 1; y >= 0; y-- {
			s = append(s, (*m)[y][x])
		}
		nm = append(nm, s)
	}
	*m = nm
}

func ct(m [][]rune) (string, int) {
	k := ""
	r := 0
	for y := range m {
		for x := range m[y] {
			k += string(m[y][x])
			if m[y][x] == 'O' {
				r += len(m) - y
			}
		}
	}
	return k, r
}

func main() {
	buf := make(map[string]int)
	hst := []int{}
	n := 1000000000
	c := 1000
	m := [][]rune{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := []rune{}
		for _, c := range scanner.Text() {
			l = append(l, c)
		}
		m = append(m, l)
	}
	for i := 0; i < c; i++ {
		for j := 0; j < 4; j++ {
			plt(&m)
			spn(&m)
		}
		k, r := ct(m)
		if j, p := buf[k]; !p {
			hst = append(hst, r)
			buf[k] = i
		} else {
			fmt.Println(hst[(n-j)%(i-j)+(j-1)])
			return
		}
	}
}
