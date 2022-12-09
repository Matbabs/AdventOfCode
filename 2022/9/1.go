package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	visited := make(map[string]bool)
	xh, yh, xt, yt, xh_last, yh_last := 0, 0, 0, 0, 0, 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		op := strings.Split(scanner.Text(), " ")
		d, _ := strconv.Atoi(op[1])
		for i := 0; i < d; i++ {
			switch op[0] {
			case "U":
				yh++
			case "D":
				yh--
			case "L":
				xh--
			case "R":
				xh++
			}
			if math.Abs(float64(yh-yt)) > 1 || math.Abs(float64(xh-xt)) > 1 {
				xt = xh_last
				yt = yh_last
				visited[fmt.Sprintf("%d,%d", xt, yt)] = true
			}
			xh_last = xh
			yh_last = yh
		}
	}
	fmt.Println(len(visited))
}
