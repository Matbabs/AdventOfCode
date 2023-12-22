package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type Cube struct {
	X, Y, Z int
}

type Brick struct {
	Start, End Cube
}

func fall(bck []Brick) int {
	changes := 0
	for i := range bck {
		a := &bck[i]
		fell := false
		for a.Start.Z > 1 {
			for j := i - 1; j >= 0; j-- {
				b := &bck[j]
				if (a.End.Z-1) >= b.Start.Z &&
					(a.Start.Z-1) <= b.End.Z &&
					a.End.X >= b.Start.X &&
					a.Start.X <= b.End.X &&
					a.End.Y >= b.Start.Y &&
					a.Start.Y <= b.End.Y {
					goto ctn
				}
			}
			if !fell {
				changes += 1
				fell = true
			}
			a.Start.Z -= 1
			a.End.Z -= 1
		}
	ctn:
	}
	return changes
}

func main() {
	res := 0
	r := regexp.MustCompile("\\d+")
	bck := []Brick{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cb := []Cube{}
		nbs := r.FindAllString(scanner.Text(), -1)
		for i := 0; i < 2; i++ {
			x, _ := strconv.Atoi(nbs[i*3])
			y, _ := strconv.Atoi(nbs[i*3+1])
			z, _ := strconv.Atoi(nbs[i*3+2])
			cb = append(cb, Cube{x, y, z})
		}
		bck = append(bck, Brick{cb[0], cb[1]})
	}
	slices.SortFunc(bck, func(a, b Brick) int {
		return a.Start.Z - b.Start.Z
	})
	fall(bck)
	tmp := make([]Brick, len(bck))
	for i := range bck {
		copy(tmp, bck)
		tmp[i] = Brick{}
		changes := fall(tmp)
		if changes == 0 {
			res++
		}
	}
	fmt.Println(res)
}
