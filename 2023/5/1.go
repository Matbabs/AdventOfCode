package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func mtl(n int, _map [][]int) int {
	for _, m := range _map {
		if n >= m[1] && n <= m[1]+m[2]-1 {
			return m[0] + n - m[1]
		}
	}
	return n
}

func main() {
	j := 0
	res := 0
	sds := []int{}
	var (
		sTs, sTf, fTw, wTl, lTt, tTh, hTl [][]int
	)
	re := regexp.MustCompile("\\d+")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			j++
		}
		ns := re.FindAllString(line, -1)
		if len(ns) > 0 {
			nbs := []int{}
			for _, s := range ns {
				n, _ := strconv.Atoi(s)
				nbs = append(nbs, n)
			}
			switch j {
			case 0:
				sds = nbs
			case 1:
				sTs = append(sTs, nbs)
			case 2:
				sTf = append(sTf, nbs)
			case 3:
				fTw = append(fTw, nbs)
			case 4:
				wTl = append(wTl, nbs)
			case 5:
				lTt = append(lTt, nbs)
			case 6:
				tTh = append(tTh, nbs)
			case 7:
				hTl = append(hTl, nbs)
			}
		}
	}
	for _, s := range sds {
		r := mtl(mtl(mtl(mtl(mtl(mtl(mtl(s, sTs), sTf), fTw), wTl), lTt), tTh), hTl)
		if res > r || res == 0 {
			res = r
		}
	}
	fmt.Println(res)
}
