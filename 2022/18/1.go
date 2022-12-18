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
	for _, drop := range drops {
		surfaces += len(freeNeighbors(drops, drop))
	}
	fmt.Println(surfaces)
}
