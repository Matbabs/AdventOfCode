package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var edg = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func dfs(_map []string, sx int, sy int) [][]int {
	st := fmt.Sprintf("%d,%d", sx, sy)
	prt := map[string]string{}
	dst := map[string]int{st: 0}
	vst := []string{st}
	tVst := [][]int{{sx, sy}}
	for len(tVst) > 0 {
		v := tVst[len(tVst)-1]
		tVst = tVst[:len(tVst)-1]
		for i, e := range edg {
			x, y := v[0]+e[0], v[1]+e[1]
			if y >= 0 && y < len(_map) && x >= 0 && x < len(_map[y]) {
				c := string(_map[y][x])
				a := false
				switch i {
				case 0:
					a = strings.Contains("J-7", c)
				case 1:
					a = strings.Contains("J|L", c)
				case 2:
					a = strings.Contains("F-L", c)
				case 3:
					a = strings.Contains("F|7", c)
				}
				if a {
					sc := fmt.Sprintf("%d,%d", x, y)
					if !slices.Contains(vst, sc) {
						vst = append(vst, sc)
						tVst = append(tVst, []int{x, y})
						p := fmt.Sprintf("%d,%d", v[0], v[1])
						dst[sc] = dst[p] + 1
						prt[sc] = p
						if dst[sc] == 2 {
							vst = []string{st, p, sc}
						}
					}
				}
			}
		}
	}
	mn := ""
	mxd := 0
	for n, d := range dst {
		if mxd < d {
			mxd = d
			mn = n
		}
	}
	l := [][]int{}
	c := mn
	for c != st {
		n := strings.Split(c, ",")
		x, _ := strconv.Atoi(n[0])
		y, _ := strconv.Atoi(n[1])
		l = append(l, []int{x, y})
		c = prt[c]
	}
	l = append(l, []int{sx, sy})
	return l
}

func rayCast(pt []int, pl [][]int) bool {
	n, inter := len(pl), 0
	for i := 0; i < n; i++ {
		p1, p2 := pl[i], pl[(i+1)%n]
		if pt[0] == p1[0] && pt[1] == p1[1] || pt[0] == p2[0] && pt[1] == p2[1] {
			return false
		}
		if (pt[1] > p1[1]) != (pt[1] > p2[1]) {
			maxX := 0
			if p1[0] > p2[0] {
				maxX = p1[0]
			} else {
				maxX = p2[0]
			}
			if pt[0] <= maxX && (p1[1] != p2[1]) {
				xInter := float64(pt[1]-p1[1])*(float64(p2[0])-float64(p1[0]))/float64(p2[1]-p1[1]) + float64(p1[0])
				if float64(p1[0]) == float64(p2[0]) || float64(pt[0]) <= xInter {
					inter++
				}
			}
		}
	}
	return inter%2 != 0
}

func main() {
	res := 0
	sy, sx := 0, 0
	_map := []string{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_map = append(_map, scanner.Text())
		for i, r := range _map[len(_map)-1] {
			if string(r) == "S" {
				sx = i
				sy = len(_map) - 1
			}
		}
	}
	l := dfs(_map, sx, sy)
	for y := range _map {
		for x := range _map[y] {
			if rayCast([]int{x, y}, l) {
				res++
			}
		}
	}
	fmt.Println(res)
}
