package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sync"
)

func bfs(m []string, sx int, sy int, d int) int {
	vst := []string{fmt.Sprintf("%d,%d,%c", sx, sy, d)}
	tVst := [][]int{{sx, sy, d}}
	eng := make(map[string]bool)
	eng[fmt.Sprintf("%d,%d", sx, sy)] = true
	for len(tVst) > 0 {
		v := tVst[0]
		tVst = tVst[1:]
		w := [][]int{}
		x, y := v[0], v[1]
		switch v[2] {
		case 'r':
			xw := x + 1
			yw1, yw2 := y-1, y+1
			switch m[y][x] {
			case '-', '.':
				if xw < len(m[0]) {
					w = append(w, []int{xw, y, 'r'})
				}
			case '|':
				if yw1 >= 0 {
					w = append(w, []int{x, yw1, 't'})
				}
				if yw2 < len(m) {
					w = append(w, []int{x, yw2, 'b'})
				}
			case '/':
				if yw1 >= 0 {
					w = append(w, []int{x, yw1, 't'})
				}
			case '\\':
				if yw2 < len(m) {
					w = append(w, []int{x, yw2, 'b'})
				}
			}
		case 'b':
			yw := y + 1
			xw1, xw2 := x-1, x+1
			switch m[y][x] {
			case '|', '.':
				if yw < len(m) {
					w = append(w, []int{x, yw, 'b'})
				}
			case '-':
				if xw1 >= 0 {
					w = append(w, []int{xw1, y, 'l'})
				}
				if xw2 < len(m[0]) {
					w = append(w, []int{xw2, y, 'r'})
				}
			case '/':
				if xw1 >= 0 {
					w = append(w, []int{xw1, y, 'l'})
				}
			case '\\':
				if xw2 < len(m[0]) {
					w = append(w, []int{xw2, y, 'r'})
				}
			}
		case 'l':
			xw := x - 1
			yw1, yw2 := y-1, y+1
			switch m[y][x] {
			case '-', '.':
				if xw >= 0 {
					w = append(w, []int{xw, y, 'l'})
				}
			case '|':
				if yw1 >= 0 {
					w = append(w, []int{x, yw1, 't'})
				}
				if yw2 < len(m) {
					w = append(w, []int{x, yw2, 'b'})
				}
			case '/':
				if yw2 < len(m) {
					w = append(w, []int{x, yw2, 'b'})
				}
			case '\\':
				if yw1 >= 0 {
					w = append(w, []int{x, yw1, 't'})
				}
			}
		case 't':
			yw := y - 1
			xw1, xw2 := x-1, x+1
			switch m[y][x] {
			case '|', '.':
				if yw >= 0 {
					w = append(w, []int{x, yw, 't'})
				}
			case '-':
				if xw1 >= 0 {
					w = append(w, []int{xw1, y, 'l'})
				}
				if xw2 < len(m[0]) {
					w = append(w, []int{xw2, y, 'r'})
				}
			case '/':
				if xw2 < len(m[0]) {
					w = append(w, []int{xw2, y, 'r'})
				}
			case '\\':
				if xw1 >= 0 {
					w = append(w, []int{xw1, y, 'l'})
				}
			}
		}
		for _, cw := range w {
			sn := fmt.Sprintf("%d,%d,%c", cw[0], cw[1], cw[2])
			if !slices.Contains(vst, sn) {
				tVst = append(tVst, []int{cw[0], cw[1], cw[2]})
				vst = append(vst, sn)
				eng[fmt.Sprintf("%d,%d", cw[0], cw[1])] = true
			}
		}

	}
	return len(eng)
}

func main() {
	r := 0
	var wg sync.WaitGroup
	res := make(chan int)
	m := []string{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m = append(m, scanner.Text())
	}
	for y := range m {
		for x := range m[y] {
			wg.Add(1)
			go func(x, y int) {
				defer wg.Done()
				if x == 0 {
					res <- bfs(m, x, y, 'r')
				}
				if x == len(m[y])-1 {
					res <- bfs(m, x, y, 'l')
				}
				if y == 0 {
					res <- bfs(m, x, y, 'b')
				}
				if y == len(m)-1 {
					res <- bfs(m, x, y, 't')
				}
			}(x, y)
		}
	}
	go func() {
		wg.Wait()
		close(res)
	}()
	for rl := range res {
		if rl > r {
			r = rl
		}
	}
	fmt.Println(r)
}
