package main

import (
	"fmt"
	"os/exec"

	"github.com/common-nighthawk/go-figure"
)

const WHITE = "\033[37m"
const PURPLE = "\033[35m"
const YELLOW = "\033[0;33m"

var excludeDays = map[int][]bool{
	17: {false, true},
	19: {true, true},
}

func displayTitle() {
	title := figure.NewColorFigure("Advent Of Code !", "", "green", true)
	title.Print()
	fmt.Printf("\n --> Matbabs - 2022\n")
}

func displayDay(day int) {
	fmt.Printf("\n%s* Day %d *\n\n", WHITE, day)
}

func displayInfo(s string) {
	fmt.Printf("%s%s", WHITE, s)
}

func executeDayByPart(day int, part int) func() {
	out, err := exec.Command("bash", "-c", fmt.Sprintf("cd %d && go run %d.go", day, part)).Output()
	if err == nil {
		color := PURPLE
		if part == 2 {
			color = YELLOW
		}
		return func() {
			fmt.Printf("%s%s\n", color, string(out)[:len(out)-1])
		}
	}
	return nil
}

func main() {
	displayTitle()
	for i := 1; i <= 25; i++ {
		for j := 1; j <= 2; j++ {
			if j == 1 {
				displayDay(i)
			}
			parts, excluded := excludeDays[i]
			if excluded && parts[j-1] {
				displayInfo("excluded for timeout\n")
				continue
			}
			displayResult := executeDayByPart(i, j)
			if displayResult != nil {
				displayResult()
			}
		}
	}
}
