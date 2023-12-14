package main

import (
	"bufio"
	"fmt"
	"os"
)

func plt(m [][]rune) []int {
	p := []int{}
	for y := range m {
		for x := range m[y] {
			if m[y][x] == 'O' {
				k := 1
				for y-k >= 0 && m[y-k][x] == '.' {
					k++
				}
				k--
				m[y][x] = '.'
				m[y-k][x] = 'O'
				p = append(p, len(m)-(y-k))
			}
		}
	}
	return p
}

func main() {
	res := 0
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
	rk := plt(m)
	for _, r := range rk {
		res += r
	}
	fmt.Println(res)
}
