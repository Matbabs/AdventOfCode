package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func hash(s string) int {
	v := 0
	for _, r := range s {
		v += int(r)
		v *= 17
		v %= 256
	}
	return v
}

func main() {
	m := make(map[int][]string)
	f := make(map[int]map[string]int)
	res := 0
	seq := []string{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seq = strings.Split(scanner.Text(), ",")
	}
	for _, s := range seq {
		if strings.Contains(s, "=") {
			c := strings.Split(s, "=")
			e, b := c[0], hash(c[0])
			n, _ := strconv.Atoi(c[1])
			if _, p := m[b]; !p {
				m[b] = []string{}
			}
			if !slices.Contains(m[b], e) {
				m[b] = append(m[b], e)
			}
			if _, p := f[b]; !p {
				f[b] = make(map[string]int)
			}
			if _, p := f[b][e]; !p {
				f[b][e] = 0
			}
			f[b][e] = n
		}
		if strings.Contains(s, "-") {
			c := strings.Split(s, "-")
			e, b, x := c[0], hash(c[0]), -1
			for i := range m[b] {
				if m[b][i] == e {
					x = i
					break
				}
			}
			if x != -1 {
				m[b] = append(m[b][:x], m[b][x+1:]...)
			}
		}
	}
	for b := range m {
		for i, e := range m[b] {
			res += (1 + b) * (1 + i) * f[b][e]
		}
	}
	fmt.Println(res)
}
