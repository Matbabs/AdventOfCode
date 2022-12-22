package main

import (
	"bufio"
	"fmt"
	"math"
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
	sideLength := int(math.Sqrt(float64(len(monkeyMap._map) / 6)))
	directionUp := ConstraintPoint[int]{R: -1, C: 0}
	directionDown := ConstraintPoint[int]{R: 1, C: 0}
	directionLeft := ConstraintPoint[int]{R: 0, C: -1}
	directionRight := ConstraintPoint[int]{R: 0, C: 1}
	nextPosDir := func(position, direction ConstraintPoint[int]) (ConstraintPoint[int], ConstraintPoint[int]) {
		newPos := ConstraintPoint[int]{
			R: position.R + direction.R,
			C: position.C + direction.C,
		}
		newDir := direction
		_, exists := monkeyMap._map[newPos]
		if !exists {
			switch direction.C {
			case 1:
				switch newPos.R / sideLength {
				case 0:
					newPos = ConstraintPoint[int]{R: sideLength*2 + (sideLength - 1 - newPos.R), C: sideLength*2 - 1}
					newDir = directionLeft
				case 1:
					newPos = ConstraintPoint[int]{R: sideLength - 1, C: newPos.R + sideLength}
					newDir = directionUp
				case 2:
					newPos = ConstraintPoint[int]{R: sideLength*3 - 1 - newPos.R, C: sideLength*3 - 1}
					newDir = directionLeft
				case 3:
					newPos = ConstraintPoint[int]{R: 3*sideLength - 1, C: newPos.R - 2*sideLength}
					newDir = directionUp
				}
			case -1:
				switch newPos.R / sideLength {
				case 0:
					newPos = ConstraintPoint[int]{R: 3*sideLength - 1 - newPos.R, C: 0}
					newDir = directionRight
				case 1:
					newPos = ConstraintPoint[int]{R: 2 * sideLength, C: newPos.R - sideLength}
					newDir = directionDown
				case 2:
					newPos = ConstraintPoint[int]{R: 3*sideLength - 1 - newPos.R, C: sideLength}
					newDir = directionRight
				case 3:
					newPos = ConstraintPoint[int]{R: 0, C: newPos.R - sideLength*2}
					newDir = directionDown
				}
			}
			switch direction.R {
			case 1:
				switch newPos.C / sideLength {
				case 0:
					newPos = ConstraintPoint[int]{R: 0, C: newPos.C + sideLength*2}
					newDir = directionDown
				case 1:
					newPos = ConstraintPoint[int]{R: newPos.C + 2*sideLength, C: sideLength - 1}
					newDir = directionLeft
				case 2:
					newPos = ConstraintPoint[int]{R: newPos.C - sideLength, C: sideLength*2 - 1}
					newDir = directionLeft
				}
			case -1:
				switch newPos.C / sideLength {
				case 0:
					newPos = ConstraintPoint[int]{R: newPos.C + sideLength, C: sideLength}
					newDir = directionRight
				case 1:
					newPos = ConstraintPoint[int]{R: newPos.C + 2*sideLength, C: 0}
					newDir = directionRight
				case 2:
					newPos = ConstraintPoint[int]{R: sideLength*4 - 1, C: newPos.C - 2*sideLength}
					newDir = directionUp
				}

			}
		}
		val, _ := monkeyMap._map[newPos]
		if val {
			return position, direction
		} else {
			return newPos, newDir
		}
	}
	position := monkeyMap.topLeft
	direction := ConstraintPoint[int]{R: 0, C: 1}
	for _, moves := range monkeyMap.moves {
		if moves.turn == "" {
			for x := 0; x < moves.magnitude; x++ {
				newPos, newDir := nextPosDir(position, direction)
				if newPos == position {
					break
				} else {
					position = newPos
					direction = newDir
				}
			}
		} else if moves.turn == RIGHT {
			direction = ConstraintPoint[int]{
				R: direction.C,
				C: -direction.R,
			}
		} else if moves.turn == LEFT {
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
