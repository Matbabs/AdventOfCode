package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ext(t []int) int {
	a := [][]int{t}
	l := []int{}
	f := true
	for f {
		l = []int{}
		b := a[len(a)-1]
		f = false
		for i := range b {
			if i+1 < len(b) {
				n := b[i+1] - b[i]
				l = append(l, n)
				if n != 0 {
					f = true
				}
			}
		}
		a = append(a, l)
	}
	for i := len(a) - 1; i >= 0; i-- {
		if i == len(a)-1 {
			a[i] = append(a[i], 0)
		} else {
			a[i] = append(a[i], a[i+1][len(a[i+1])-1]+a[i][len(a[i])-1])
		}
	}
	return a[0][len(a[0])-1]
}

func main() {
	res := 0
	h := [][]int{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := []int{}
		nss := strings.Split(scanner.Text(), " ")
		for _, ns := range nss {
			n, _ := strconv.Atoi(ns)
			l = append(l, n)
		}
		h = append(h, l)
	}
	for _, t := range h {
		res += ext(t)
	}
	fmt.Println(res)
}
