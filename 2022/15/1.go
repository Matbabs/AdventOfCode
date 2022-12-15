package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

const ROW = 2000000

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

func main() {
	score := 0
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
	for x := min_x; x <= max_x; x++ {
		for _, c := range pairs {
			if x == c[2] && ROW == c[3] {
				break
			}
			if manhattanDist(x, c[0], ROW, c[1]) <= c[len(c)-1] {
				score++
				break
			}
		}
	}
	fmt.Println(score)
}
