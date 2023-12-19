package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const ST = 1
const ED = 4000

var CONF = []string{"x", "m", "a", "s"}

func prc(pc map[string][]int, swkn string, wkfs map[string][]string, acc *[]map[string][]int) {
	wkf := wkfs[swkn]
	for _, c := range wkf {
		if strings.Contains(c, ":") {
			cs := strings.Split(c, ":")
			pcn := string(cs[0][0])
			pcc := cs[0][1]
			pcv, _ := strconv.Atoi(cs[0][2:])
			ewkn := cs[1]
			a := make(map[string][]int)
			b := make(map[string][]int)
			for _, l := range CONF {
				a[l] = []int{pc[l][0], pc[l][1]}
				b[l] = []int{pc[l][0], pc[l][1]}
			}
			if pcc == '<' {
				a[pcn][0] = pcv
				b[pcn][1] = pcv - 1
				prc(b, ewkn, wkfs, acc)
				pc = a
			}
			if pcc == '>' {
				a[pcn][0] = pcv + 1
				b[pcn][1] = pcv
				prc(a, ewkn, wkfs, acc)
				pc = b
			}
		} else {
			if c == "A" {
				*acc = append(*acc, pc)
			} else {
				prc(pc, c, wkfs, acc)
			}
		}
	}
}

func main() {
	res := int64(0)
	lb := 0
	wkfs := map[string][]string{"A": {"A"}, "E": {"E"}}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			lb++
		} else if lb == 0 {
			wfs := strings.Split(l, "{")
			wfn := wfs[0]
			wfp := strings.Split(strings.Split(wfs[1], "}")[0], ",")
			wkfs[wfn] = wfp
		}
	}
	var rng []map[string][]int
	st := make(map[string][]int)
	for _, l := range CONF {
		st[l] = []int{ST, ED}
	}
	prc(st, "in", wkfs, &rng)
	for _, r := range rng {
		rl := int64(1)
		for _, l := range CONF {
			rl *= int64(r[l][1] - r[l][0] + 1)
		}
		res += rl
	}
	fmt.Println(res)
}
