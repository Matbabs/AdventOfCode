package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

const MIN_INT = 0
const MAX_INT = 4000000
const ROW = 2000000

var TURNS = []struct{ sin, cos int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func convertStringsInts(a []string) []int {
	res := make([]int, 0)
	for _, n := range a {
		v, _ := strconv.Atoi(n)
		res = append(res, v)
	}
	return res
}

func manhattanDist(x1, x2, y1, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}

func diamond(x, y, dist int) [][]int {
	res := make([][]int, 0)
	for _, turn := range TURNS {
		for dx := 0; dx <= dist; dx++ {
			dy := dist - dx
			res = append(res, []int{x + dx*turn.cos - dy*turn.sin, y + dx*turn.sin + dy*turn.cos})
		}
	}
	return res
}

func main() {
	min_x, max_x := -1, 0
	pairs := make([][]int, 0)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`-?\d+`)
		cds := convertStringsInts(re.FindAllString(line, -1))
		xs, ys, xb, yb := cds[0], cds[1], cds[2], cds[3]
		dist := manhattanDist(xs, xb, ys, yb)
		l, r := xs-dist, xs+dist
		if min_x == -1 || min_x > l {
			min_x = l
		}
		if max_x < r {
			max_x = r
		}
		pairs = append(pairs, []int{xs, ys, xb, yb, dist})
	}
	for {
		for _, c := range pairs {
			for _, cd := range diamond(c[0], c[1], c[len(c)-1]+1) {
				if cd[0] >= MIN_INT && cd[0] <= MAX_INT && cd[1] >= MIN_INT && cd[1] <= MAX_INT {
					isValid := true
					for _, c2 := range pairs {
						if manhattanDist(cd[0], c2[0], cd[1], c2[1]) <= c2[len(c)-1] {
							isValid = false
							break
						}
					}
					if isValid {
						fmt.Println(cd[0]*MAX_INT + cd[1])
						return
					}
				}
			}
		}
	}
}
