package main

import (
	"bufio"
	"fmt"
	"os"
)

const STEPS = 64

var drs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func bfs(m []string, sx, sy, steps int) map[string]int {
	dst := map[string]int{fmt.Sprintf("%d,%d", sx, sy): 0}
	tVst := [][]int{{sx, sy, steps}}
	for len(tVst) > 0 {
		v := tVst[0]
		tVst = tVst[1:]
		for _, d := range drs {
			wx, wy := v[0]+d[0], v[1]+d[1]
			if wy >= 0 && wy < len(m) && wx >= 0 && wx < len(m[wy]) && m[wy][wx] != '#' {
				sw := fmt.Sprintf("%d,%d", wx, wy)
				if _, p := dst[sw]; !p && v[2]-1 >= 0 {
					tVst = append(tVst, []int{wx, wy, v[2] - 1})
					dst[sw] = dst[fmt.Sprintf("%d,%d", v[0], v[1])] + 1
				}
			}
		}
	}
	return dst
}

func main() {
	res := 0
	sx, sy := -1, 0
	m := []string{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m = append(m, scanner.Text())
		for i, r := range m[len(m)-1] {
			if r == 'S' {
				sx = i
			}
		}
		if sx == -1 {
			sy++
		}
	}
	for _, d := range bfs(m, sx, sy, STEPS) {
		if (d+STEPS%2)%2 == 0 {
			res++
		}
	}
	fmt.Println(res)
}
