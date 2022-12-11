package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const TURNS = 10000

type Monkey struct {
	items []int
	ops   string
	test  int
	true  int
	false int
}

func getInt(line string) int {
	split := strings.Split(line, " ")
	v, _ := strconv.Atoi(split[len(split)-1])
	return v
}

func main() {
	inspections := make([]int, 0)
	monkeys := make([]Monkey, 0)
	monkey := Monkey{items: []int{}}
	num_line := 0
	solve := 1
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
			monkey.test = getInt(line)
			solve *= monkey.test
		case 4:
			monkey.true = getInt(line)
		case 5:
			monkey.false = getInt(line)
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
				cmd, _ := exec.Command("bash", "-c", fmt.Sprintf("echo $((%s))", exp)).Output()
				lvl, _ := strconv.Atoi(string(cmd[:len(cmd)-1]))
				lvl = lvl % solve
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
