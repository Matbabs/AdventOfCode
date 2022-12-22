package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"golang.org/x/exp/constraints"
)

type Rotation string

const (
	LEFT       Rotation = "L"
	RIGHT      Rotation = "R"
	ROW_FACTOR int      = 1000
	COL_FACTOR int      = 4
)

type ConstraintPoint[T constraints.Signed] struct {
	R, C T
}

type Move struct {
	magnitude int
	turn      Rotation
}

type MonkeyMap struct {
	_map    map[ConstraintPoint[int]]bool
	topLeft ConstraintPoint[int]
	maxC    int
	maxR    int
	moves   []Move
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	monkeyMap := MonkeyMap{}
	monkeyMap._map = map[ConstraintPoint[int]]bool{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for r, line := range lines {
		if line == "" {
			re := regexp.MustCompile(`(?:(\d+)|(L|R))`)
			submatch := re.FindAllStringSubmatch(lines[r+1], -1)
			for _, s := range submatch {
				if s[1] != "" {
					v, _ := strconv.Atoi(s[1])
					monkeyMap.moves = append(monkeyMap.moves, Move{
						magnitude: v,
					})
				} else {
					monkeyMap.moves = append(monkeyMap.moves, Move{
						turn: Rotation(s[2]),
					})
				}
			}
			break
		} else {
			found := false
			for c, char := range line {
				switch char {
				case '.':
					monkeyMap._map[ConstraintPoint[int]{R: r, C: c}] = false
					if r == 0 && !found {
						monkeyMap.topLeft = ConstraintPoint[int]{R: r, C: c}
					}
					found = true
				case '#':
					monkeyMap._map[ConstraintPoint[int]{R: r, C: c}] = true
				}
				monkeyMap.maxC = max(monkeyMap.maxC, c)
			}
			monkeyMap.maxR = max(monkeyMap.maxR, r)
		}
	}
	position := monkeyMap.topLeft
	direction := ConstraintPoint[int]{R: 0, C: 1}
	for _, move := range monkeyMap.moves {
		if move.turn == "" {
			for x := 0; x < move.magnitude; x++ {
				newPos := ConstraintPoint[int]{
					R: position.R + direction.R,
					C: position.C + direction.C,
				}
				_, exists := monkeyMap._map[newPos]
				if !exists {
					if direction.C == 1 {
						for c := 0; ; c++ {
							if _, exists := monkeyMap._map[ConstraintPoint[int]{R: newPos.R, C: c}]; exists {
								newPos = ConstraintPoint[int]{R: newPos.R, C: c}
								break
							}
						}
					} else if direction.C == -1 {
						for c := monkeyMap.maxC; ; c-- {
							if _, exists := monkeyMap._map[ConstraintPoint[int]{R: newPos.R, C: c}]; exists {
								newPos = ConstraintPoint[int]{R: newPos.R, C: c}
								break
							}
						}
					} else if direction.R == 1 {
						for r := 0; ; r++ {
							if _, exists := monkeyMap._map[ConstraintPoint[int]{R: r, C: newPos.C}]; exists {
								newPos = ConstraintPoint[int]{R: r, C: newPos.C}
								break
							}
						}
					} else if direction.R == -1 {
						for r := monkeyMap.maxR; ; r-- {
							if _, exists := monkeyMap._map[ConstraintPoint[int]{R: r, C: newPos.C}]; exists {
								newPos = ConstraintPoint[int]{R: r, C: newPos.C}
								break
							}
						}
					}
				}
				if wall, exists := monkeyMap._map[newPos]; exists && wall {
					break
				} else {
					position = newPos
				}
			}
		} else if move.turn == RIGHT {
			direction = ConstraintPoint[int]{
				R: direction.C,
				C: -direction.R,
			}
		} else if move.turn == LEFT {
			direction = ConstraintPoint[int]{
				R: -direction.C,
				C: direction.R,
			}
		}
	}
	face := 0
	if direction.C == 1 {
		face = 0
	} else if direction.C == -1 {
		face = 2
	} else if direction.R == 1 {
		face = 1
	} else {
		face = 3
	}
	fmt.Println((position.R+1)*ROW_FACTOR + (position.C+1)*COL_FACTOR + face)
}
