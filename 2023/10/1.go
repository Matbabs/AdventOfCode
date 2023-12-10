package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var edg = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func bfs(_map []string, sx int, sy int) map[string]int {
	dst := map[string]int{fmt.Sprintf("%d,%d", sx, sy): 0}
	vst := []string{fmt.Sprintf("%d,%d", sx, sy)}
	tVst := [][]int{{sx, sy}}
	for len(tVst) > 0 {
		v := tVst[0]
		tVst = tVst[1:]
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
						dst[sc] = dst[fmt.Sprintf("%d,%d", v[0], v[1])] + 1
					}
				}
			}
		}
	}
	return dst
}

func main() {
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
	res := 0
	for _, d := range bfs(_map, sx, sy) {
		if d > res {
			res = d
		}
	}
	fmt.Println(res)
}
