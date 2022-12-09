package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const NB_NODES = 10

func main() {
	visited := make(map[string]bool)
	nodes := make([][]int, NB_NODES)
	lasts := make([][]int, NB_NODES)
	for i := 0; i < NB_NODES; i++ {
		nodes[i] = make([]int, 2)
		lasts[i] = make([]int, 2)
	}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		op := strings.Split(scanner.Text(), " ")
		d, _ := strconv.Atoi(op[1])
		for i := 0; i < d; i++ {
			switch op[0] {
			case "U":
				nodes[0][1] += 1
			case "D":
				nodes[0][1] -= 1
			case "L":
				nodes[0][0] -= 1
			case "R":
				nodes[0][0] += 1
			}
			for j := range nodes[1:] {
				hx, hy := nodes[j][0], nodes[j][1]
				tx, ty := nodes[j+1][0], nodes[j+1][1]
				dx := math.Abs(float64(hx - tx))
				dy := math.Abs(float64(hy - ty))
				if dx > 1 {
					if hx > tx {
						tx++
					} else {
						tx--
					}
					if dy > 0 {
						if hy > ty {
							ty++
						} else {
							ty--
						}
					}
				} else if dy > 1 {
					if hy > ty {
						ty++
					} else {
						ty--
					}
					if dx > 0 {
						if hx > tx {
							tx++
						} else {
							tx--
						}
					}
				}
				nodes[j+1][0] = tx
				nodes[j+1][1] = ty
				if j == NB_NODES-2 {
					visited[fmt.Sprintf("%d,%d", tx, ty)] = true
				}
			}
		}
	}
	fmt.Println(len(visited))
}
