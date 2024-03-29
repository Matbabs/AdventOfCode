package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var drs = map[string][2]int{
	"R": {0, 1},
	"D": {1, 0},
	"L": {0, -1},
	"U": {-1, 0},
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func tupleAdd(t1, t2 [2]int, multiply int) [2]int {
	return [2]int{t1[0] + t2[0]*multiply, t1[1] + t2[1]*multiply}
}

func main() {
	var plan [][]interface{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		dir := drs[fields[0]]
		mts, _ := strconv.Atoi(fields[1])
		plan = append(plan, []interface{}{dir, mts})
	}
	cds := [][2]int{{0, 0}}
	pPos := [2]int{0, 0}
	bdl := 1
	for _, entry := range plan {
		dir := entry[0].([2]int)
		mts := entry[1].(int)
		nPos := tupleAdd(pPos, dir, mts)
		bdl += mts
		cds = append(cds, nPos)
		pPos = nPos
	}
	area := 0
	for i := 1; i < len(cds); i++ {
		x1, y1 := cds[i-1][0], cds[i-1][1]
		x2, y2 := cds[i][0], cds[i][1]
		area += (y1 + y2) * (x1 - x2)
	}
	area = (abs(area) + bdl + 1) / 2
	fmt.Println(area)
}
