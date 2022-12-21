package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Operation string

const (
	ROOT  string    = "root"
	PLUS  Operation = "+"
	MINUS Operation = "-"
	DIV   Operation = "/"
	MUL   Operation = "*"
)

type Monkey struct {
	name        string
	operation   Operation
	left, right string
	value       int
}

func solve(m Monkey, monkeys map[string]Monkey) int {
	if m.value > 0 {
		return m.value
	}
	l := solve(monkeys[m.left], monkeys)
	r := solve(monkeys[m.right], monkeys)
	switch m.operation {
	case PLUS:
		return l + r
	case MINUS:
		return l - r
	case DIV:
		return l / r
	case MUL:
		return l * r
	default:
		return -1
	}
}

func main() {
	monkeys := make(map[string]Monkey)
	re := regexp.MustCompile(`(\w+) (.) (\w+)`)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ": ")
		m := Monkey{name: split[0]}
		val, err := strconv.Atoi(split[1])
		if err == nil {
			m.value = int(val)
		} else {
			op := re.FindStringSubmatch(split[1])[1:]
			m.left = op[0]
			m.operation = Operation(op[1])
			m.right = op[2]
		}
		monkeys[split[0]] = m
	}
	result := solve(monkeys[ROOT], monkeys)
	fmt.Println(result)
}
