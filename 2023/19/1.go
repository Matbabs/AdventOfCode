package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func prc(pc map[string]int, swkn string, wkfs map[string][]string) int {
	wkf := wkfs[swkn]
	for _, c := range wkf {
		if strings.Contains(c, ":") {
			cs := strings.Split(c, ":")
			pcn := string(cs[0][0])
			pcc := cs[0][1]
			pcv, _ := strconv.Atoi(cs[0][2:])
			ewkn := cs[1]
			if pcc == '>' && pc[pcn] > pcv {
				return prc(pc, ewkn, wkfs)
			}
			if pcc == '<' && pc[pcn] < pcv {
				return prc(pc, ewkn, wkfs)
			}
		} else {
			switch c {
			case "A":
				sm := 0
				for _, v := range pc {
					sm += v
				}
				return sm
			case "E":
				return 0
			default:
				return prc(pc, c, wkfs)
			}
		}
	}
	return 0
}

func main() {
	r, lb := 0, 0
	res := make(chan int)
	var wg sync.WaitGroup
	wkfs := map[string][]string{"A": {"A"}, "E": {"E"}}
	pcs := []map[string]int{}
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
		} else {
			p := map[string]int{}
			ps := strings.Split(strings.Split(strings.Split(l, "{")[1], "}")[0], ",")
			for _, pv := range ps {
				pvs := strings.Split(pv, "=")
				pvn, _ := strconv.Atoi(pvs[1])
				p[pvs[0]] = pvn
			}
			pcs = append(pcs, p)
		}
	}
	for i := range pcs {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			res <- prc(pcs[i], "in", wkfs)
		}(i)
	}
	go func() {
		wg.Wait()
		close(res)
	}()
	for rl := range res {
		r += rl
	}
	fmt.Println(r)
}
