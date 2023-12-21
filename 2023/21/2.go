package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const STEPS = 26501365

var drs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func bfs(m []string, sx, sy, steps int) map[string]int {
	dst := map[string]int{fmt.Sprintf("%d,%d", sx, sy): 0}
	tVst := [][]int{{sx, sy, steps}}
	for len(tVst) > 0 {
		v := tVst[0]
		tVst = tVst[1:]
		for _, d := range drs {
			wx, wy := v[0]+d[0], v[1]+d[1]
			tcwx, tcwy := wx, wy
			if wy >= len(m) {
				tcwy = wy % len(m)
			}
			if wy < 0 {
				tcwy = (wy%len(m) + len(m)) % len(m)
			}
			if wx >= len(m[tcwy]) {
				tcwx = wx % len(m[tcwy])
			}
			if wx < 0 {
				tcwx = (wx%len(m[tcwy]) + len(m[tcwy])) % len(m[tcwy])
			}
			if m[tcwy][tcwx] != '#' {
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

func fstTermsAr(n int64, p ...int64) int64 {
	return p[0] + n*(p[1]-p[0]) + n*(n-1)/2*((p[2]-p[1])-(p[1]-p[0]))
}

func main() {
	sx, sy := -1, 0
	m := []string{}
	params := []int64{}
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
	for i := 0; i < len(m)*3; i++ {
		if i%len(m) == int(math.Floor(float64(len(m))/float64(2))) {
			r := 0
			for _, d := range bfs(m, sx, sy, i) {
				if (d+i%2)%2 == 0 {
					r++
				}
			}
			params = append(params, int64(r))
		}
	}
	fmt.Println(fstTermsAr(int64(math.Floor(float64(STEPS)/float64(len(m)))), params...))
}
