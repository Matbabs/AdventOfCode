package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Drop [3]int

func contains(drops []Drop, drop Drop) bool {
	for _, c := range drops {
		if c == drop {
			return true
		}
	}
	return false
}

func freeNeighbors(drops []Drop, drop Drop) []Drop {
	var neighbors []Drop
	for i := 0; i < 3; i++ {
		for d := -1; d <= 1; d += 2 {
			p := drop
			p[i] += d
			if !contains(drops, p) {
				neighbors = append(neighbors, p)
			}
		}
	}
	return neighbors
}

func floodfill(drops []Drop, drop Drop, pockets, outside map[Drop]bool, mins, maxs [3]int) bool {
	toProceed := []Drop{drop}
	connected := make(map[Drop]bool)
	for len(toProceed) > 0 {
		d := toProceed[len(toProceed)-1]
		toProceed = toProceed[:len(toProceed)-1]
		for _, n := range freeNeighbors(drops, d) {
			if n[0] < mins[0] || n[0] > maxs[0] || n[1] < mins[1] || n[1] > maxs[1] || n[2] < mins[2] || n[2] > maxs[2] {
				outside[d] = true
				return true
			}
			if _, ok := connected[n]; !ok {
				connected[n] = true
				toProceed = append(toProceed, n)
			}
		}
	}
	for coord := range connected {
		pockets[coord] = true
	}
	return false
}

func main() {
	surfaces := 0
	drops := make([]Drop, 0)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := make([]int, 0)
		re := regexp.MustCompile(`\d+`)
		for _, n := range re.FindAllString(scanner.Text(), -1) {
			v, _ := strconv.Atoi(n)
			c = append(c, v)
		}
		drops = append(drops, Drop{c[0], c[1], c[2]})
	}
	mins := [3]int{drops[0][0], drops[0][1], drops[0][2]}
	maxs := [3]int{drops[0][0], drops[0][1], drops[0][2]}
	for _, drop := range drops {
		if drop[0] < mins[0] {
			mins[0] = drop[0]
		} else if drop[0] > maxs[0] {
			maxs[0] = drop[0]
		}
		if drop[1] < mins[1] {
			mins[1] = drop[1]
		} else if drop[1] > maxs[1] {
			maxs[1] = drop[1]
		}
		if drop[2] < mins[2] {
			mins[2] = drop[2]
		} else if drop[2] > maxs[2] {
			maxs[2] = drop[2]
		}
	}
	pockets := make(map[Drop]bool)
	outside := make(map[Drop]bool)
	for _, drop := range drops {
		for _, neighbor := range freeNeighbors(drops, drop) {
			_, ok := pockets[neighbor]
			if !ok && (outside[neighbor] || floodfill(drops, neighbor, pockets, outside, mins, maxs)) {
				surfaces++
			}
		}
	}
	fmt.Println(surfaces)
}
