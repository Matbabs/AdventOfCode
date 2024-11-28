package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func intersectingPaths(o1, o2 []float64, max, min float64) bool {
	// Check if the velocities are equal, indicating parallel trajectories
	if o1[3] == o2[3] && o1[4] == o2[4] {
		return false
	}

	// Check if the paths will cross inside the test area
	time := (o2[0] - o1[0]) / (o1[3] - o2[3])

	// Check if time is finite and within the valid range
	if math.IsInf(time, 0) {
		return false
	}

	x1, y1 := o1[0]+o1[3]*time, o1[1]+o1[4]*time
	x2, y2 := o2[0]+o2[3]*time, o2[1]+o2[4]*time

	fmt.Println(x1, y1)

	// Check if the paths intersect inside the test area
	if x1 >= min && x1 <= max && y1 >= min && y1 <= max &&
		x2 >= min && x2 <= max && y2 >= min && y2 <= max {
		fmt.Println("Paths intersect at:", x1, y1)
		return true
	}

	return false
}

func main() {
	res := 0
	hls := [][]float64{}
	r := regexp.MustCompile("-?\\d+")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cds := []float64{}
		for _, ns := range r.FindAllString(scanner.Text(), -1) {
			n, _ := strconv.Atoi(ns)
			cds = append(cds, float64(n))
		}
		hls = append(hls, cds)
	}
	fmt.Println(hls)

	for i := 0; i < len(hls); i++ {
		for j := i + 1; j < len(hls); j++ {
			if intersectingPaths(hls[i], hls[j], 7, 27) {
				res++
			}
		}
	}
	fmt.Println(res)
}
