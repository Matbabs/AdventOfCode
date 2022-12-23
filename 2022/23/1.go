package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/constraints"
)

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Contains(val T) (contains bool) {
	_, contains = s[val]
	return
}

func (s Set[T]) List() []T {
	result := []T{}
	for x := range s {
		result = append(result, x)
	}
	return result
}

func (s Set[T]) Equal(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for k := range s {
		if !other.Contains(k) {
			return false
		}
	}
	return true
}

type ConstraintPoint[T constraints.Signed] struct {
	R, C T
}

func newConstraintPoint[T constraints.Signed](r, c T) ConstraintPoint[T] {
	return ConstraintPoint[T]{R: r, C: c}
}

func Max[T constraints.Ordered](a, b T) T {
	if a >= b {
		return a
	} else {
		return b
	}
}

func Min[T constraints.Ordered](a, b T) T {
	if a <= b {
		return a
	} else {
		return b
	}
}

func minMaxListFn[T any, U constraints.Ordered](l []T, f func(T) U) (min, max U) {
	if len(l) == 0 {
		panic("cannot find min/max of empty list")
	}
	min = f(l[0])
	max = min
	for _, val := range l {
		v := f(val)
		min = Min(min, v)
		max = Max(max, v)
	}
	return
}

type ElvesMap struct {
	_map Set[ConstraintPoint[int]]
}

func (elvesMap *ElvesMap) Round(i int, poses Set[ConstraintPoint[int]]) Set[ConstraintPoint[int]] {
	proposals := map[ConstraintPoint[int]]ConstraintPoint[int]{}
	newPoses := Set[ConstraintPoint[int]]{}
	for pose := range poses {
		surrounds := false
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				if dr == 0 && dc == 0 {
					continue
				}
				if poses.Contains(newConstraintPoint(pose.R+dr, pose.C+dc)) {
					surrounds = true
					break
				}
			}
		}
		if !surrounds {
			newPoses.Add(pose)
			continue
		}
		considers := []struct {
			dr  []int
			dc  []int
			dir ConstraintPoint[int]
		}{
			{[]int{-1}, []int{-1, 0, 1}, newConstraintPoint(-1, 0)}, // North
			{[]int{1}, []int{-1, 0, 1}, newConstraintPoint(1, 0)},   // South
			{[]int{-1, 0, 1}, []int{-1}, newConstraintPoint(0, -1)}, // West
			{[]int{-1, 0, 1}, []int{1}, newConstraintPoint(0, 1)},   // East
		}
		proposed := false
		for j := 0; j < 4; j++ {
			consider := considers[(j+i)%4]
			open := true
			for _, dr := range consider.dr {
				for _, dc := range consider.dc {
					if poses.Contains(newConstraintPoint(pose.R+dr, pose.C+dc)) {
						open = false
						break
					}
				}
			}
			if open {
				proposals[pose] = newConstraintPoint(pose.R+consider.dir.R, pose.C+consider.dir.C)
				proposed = true
				break
			}
		}
		if !proposed {
			newPoses.Add(pose)
		}
	}
	overlaps := map[ConstraintPoint[int]]int{}
	for _, proposal := range proposals {
		overlaps[proposal]++
	}
	for elf, proposal := range proposals {
		if overlaps[proposal] == 1 {
			newPoses.Add(proposal)
		} else {
			newPoses.Add(elf)
		}
	}
	return newPoses
}

func (elvesMap *ElvesMap) Part2(isTest bool) int {
	poses := elvesMap._map
	for i := 0; ; i++ {
		newPoses := elvesMap.Round(i, poses)
		if poses.Equal(newPoses) {
			return i + 1
		}
		poses = newPoses
	}
}

func main() {
	result := 0
	elvesMap := ElvesMap{}
	elvesMap._map = Set[ConstraintPoint[int]]{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for r, line := range lines {
		for c, char := range line {
			if char == '#' {
				elvesMap._map.Add(newConstraintPoint(r, c))
			}
		}
	}
	poses := elvesMap._map
	for i := 0; i < 10; i++ {
		newPoses := elvesMap.Round(i, poses)
		if poses.Equal(newPoses) {
			break
		}
		poses = newPoses
	}
	minR, maxR := minMaxListFn(poses.List(), func(p ConstraintPoint[int]) int { return p.R })
	minC, maxC := minMaxListFn(poses.List(), func(p ConstraintPoint[int]) int { return p.C })
	for r := minR; r <= maxR; r++ {
		for c := minC; c <= maxC; c++ {
			if !poses.Contains(newConstraintPoint(r, c)) {
				result++
			}
		}
	}
	fmt.Println(result)
}
