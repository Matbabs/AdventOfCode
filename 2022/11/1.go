package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const TURNS = 20

type Monkey struct {
	items []int
	ops   string
	test  int
	true  int
	false int
}

func getLastInt(line string) int {
	split := strings.Split(line, " ")
	v, _ := strconv.Atoi(split[len(split)-1])
	return v
}

func main() {
	inspections := make([]int, 0)
	monkeys := make([]Monkey, 0)
	monkey := Monkey{items: []int{}}
	num_line := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch num_line {
		case 1:
			re := regexp.MustCompile(`\d+`)
			for _, n := range re.FindAllString(line, -1) {
				v, _ := strconv.Atoi(n)
				monkey.items = append(monkey.items, v)
			}
		case 2:
			monkey.ops = strings.Split(line, "=")[1]
		case 3:
			monkey.test = getLastInt(line)
		case 4:
			monkey.true = getLastInt(line)
		case 5:
			monkey.false = getLastInt(line)
			monkeys = append(monkeys, monkey)
			inspections = append(inspections, 0)
			monkey = Monkey{items: []int{}}
			num_line = -2
		}
		num_line++
	}
	for i := 0; i < TURNS; i++ {
		for j := range monkeys {
			for len(monkeys[j].items) > 0 {
				item := monkeys[j].items[0]
				monkeys[j].items = monkeys[j].items[1:]
				exp := strings.ReplaceAll(monkeys[j].ops, "old", strconv.Itoa(item))
				ops := strings.Split(exp, " ")
				lvl, _ := strconv.Atoi(ops[1])
				opr, _ := strconv.Atoi(ops[3])
				if ops[2] == "+" {
					lvl = int(math.Floor(float64((lvl + opr) / 3)))
				} else {
					lvl = int(math.Floor(float64((lvl * opr) / 3)))
				}
				if lvl%monkeys[j].test == 0 {
					monkeys[monkeys[j].true].items = append(monkeys[monkeys[j].true].items, lvl)
				} else {
					monkeys[monkeys[j].false].items = append(monkeys[monkeys[j].false].items, lvl)
				}
				inspections[j]++
			}
		}
	}
	sort.Ints(inspections)
	fmt.Println(inspections[len(inspections)-1] * inspections[len(inspections)-2])
}
