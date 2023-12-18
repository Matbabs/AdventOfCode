package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const CRUCIBLE = 3
const INF = 10000000

var dir = map[string][]int{"r": {1, 0}, "b": {0, 1}, "l": {0, -1}, "t": {-1, 0}}

func isValidMove(m []string, a string, ax, ay int, ad, bd string, p map[string]string) bool {
	ps := []string{}
	for i := 0; i < CRUCIBLE; i++ {
		pd := strings.Split(a, ",")
		if len(pd) > 2 {
			ps = append(ps, pd[2])
			a = p[a]
		}
	}
	sd := true
	if len(ps) == CRUCIBLE {
		for _, pk := range ps {
			if bd != pk {
				sd = false
			}
		}
		if sd {
			return false
		}
	}
	switch ad {
	case "r":
		return bd != "t"
	case "b":
		return bd != "l"
	case "t":
		return bd != "r"
	case "l":
		return bd != "b"
	}
	return false
}

func dijkstra(m []string) {
	P := make(map[string]bool)
	d := make(map[string]int)
	p := make(map[string]string)
	fstSn := []string{}
	var drs []string
	for dir := range dir {
		drs = append(drs, dir)
	}
	sort.Strings(drs)
	for y := range m {
		for x := range m[y] {
			for _, dk := range drs {
				d[fmt.Sprintf("%d,%d,%s", x, y, dk)] = INF
			}
		}
	}
	st := fmt.Sprintf("%d,%d,%s", 0, 0, "r")
	d[st] = 0
	fstSn = append(fstSn, st)
	for {
		a := ""
		da := INF
		for n, dist := range d {
			if _, p := P[n]; !p && dist < da {
				da = dist
			}
		}
		pos := []string{}
		for n := range d {
			if _, p := P[n]; !p {
				if d[n] == da {
					pos = append(pos, n)
				}
			}
		}
		for i, n := range fstSn {
			if slices.Contains(pos, n) {
				a = n
				fstSn = append(fstSn[:i], fstSn[i+1:]...)
				break
			}
		}
		if da == INF {
			break
		}
		P[a] = true
		as := strings.Split(a, ",")
		ax, _ := strconv.Atoi(as[0])
		ay, _ := strconv.Atoi(as[1])
		ad := as[2]
		for _, bd := range drs {
			dr := dir[bd]
			if isValidMove(m, a, ax, ay, ad, bd, p) {
				bx, by := ax+dr[0], ay+dr[1]
				if bx >= 0 && bx < len(m[0]) && by >= 0 && by < len(m) {
					b := fmt.Sprintf("%d,%d,%s", bx, by, bd)
					w, _ := strconv.Atoi(string(m[by][bx]))
					if d[b] > d[a]+w {
						d[b] = d[a] + w
						p[b] = a
						fstSn = append(fstSn, b)
					}
				}
			}
		}
	}
	mhl := INF
	for dk := range dir {
		k := fmt.Sprintf("%d,%d,%s", len(m[0])-1, len(m)-1, dk)
		hm := d[k]
		if hm < mhl {
			mhl = hm
		}
	}
	fmt.Println(mhl)
}

func main() {
	m := []string{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m = append(m, scanner.Text())
	}
	dijkstra(m)
}
