package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
)

func mtl(n int, _map [][]int) int {
	for _, m := range _map {
		if n >= m[1] && n <= m[1]+m[2]-1 {
			return m[0] + n - m[1]
		}
	}
	return n
}

func cpt(wg *sync.WaitGroup, j int, sds []int, sTs, sTf, fTw, wTl, lTt, tTh, hTl [][]int, resChan chan int) {
	defer wg.Done()
	lr := 0
	for i := 0; i < len(sds); i += 2 {
		for k := sds[i]; k < sds[i]+sds[i+1]; k++ {
			r := mtl(mtl(mtl(mtl(mtl(mtl(mtl(k, sTs), sTf), fTw), wTl), lTt), tTh), hTl)
			if lr > r || lr == 0 {
				lr = r
			}
		}
	}
	resChan <- lr
}

func main() {
	j := 0
	res := 0
	sds := []int{}
	var (
		sTs, sTf, fTw, wTl, lTt, tTh, hTl [][]int
	)
	var wg sync.WaitGroup
	resChan := make(chan int, len(sds)/2)
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
	for i := 0; i < len(sds); i += 2 {
		wg.Add(1)
		go cpt(&wg, i, sds, sTs, sTf, fTw, wTl, lTt, tTh, hTl, resChan)
	}
	go func() {
		wg.Wait()
		close(resChan)
	}()
	for lr := range resChan {
		if res > lr || res == 0 {
			res = lr
		}
	}
	fmt.Println(res)
}
