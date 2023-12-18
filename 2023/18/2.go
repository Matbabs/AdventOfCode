package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

var drs = map[string][2]int{
	"0": {0, 1},
	"1": {1, 0},
	"2": {0, -1},
	"3": {-1, 0},
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

func hexToNb(hxs string) *big.Int {
	hxs = strings.TrimPrefix(hxs, "#")
	r, _ := new(big.Int).SetString(hxs, 16)
	return r
}

func main() {
	var plan [][]interface{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		dir := drs[string(fields[2][len(fields[2])-2:len(fields[2])-1])]
		mts := hexToNb(fields[2][1 : len(fields[2])-2])
		plan = append(plan, []interface{}{dir, mts})
	}
	cds := [][2]int{{0, 0}}
	pPos := [2]int{0, 0}
	bdl := big.NewInt(1)
	for _, entry := range plan {
		dir := entry[0].([2]int)
		mts := entry[1].(*big.Int)
		nPos := tupleAdd(pPos, dir, int(mts.Int64()))
		bdl.Add(bdl, mts)
		cds = append(cds, nPos)
		pPos = nPos
	}
	area := big.NewInt(0)
	for i := 1; i < len(cds); i++ {
		x1, y1 := cds[i-1][0], cds[i-1][1]
		x2, y2 := cds[i][0], cds[i][1]
		term := big.NewInt(0).SetInt64(int64(y1+y2) * int64(x1-x2))
		area.Add(area, term)
	}
	areaAbs := big.NewInt(0).Abs(area)
	areaAbs.Add(areaAbs, bdl)
	areaAbs.Add(areaAbs, big.NewInt(1))
	areaAbs.Div(areaAbs, big.NewInt(2))
	fmt.Println(areaAbs)
}
