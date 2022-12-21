package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Operation string

const (
	HUMAN string    = "humn"
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

func (m Monkey) lhs(closed string) Monkey {
	switch m.operation {
	case PLUS:
		return Monkey{operation: MINUS, left: closed, right: m.right}
	case MINUS:
		return Monkey{operation: PLUS, left: closed, right: m.right}
	case DIV:
		return Monkey{operation: MUL, left: closed, right: m.right}
	case MUL:
		return Monkey{operation: DIV, left: closed, right: m.right}
	default:
		return Monkey{}
	}
}

func (m Monkey) rhs(closed string) Monkey {
	switch m.operation {
	case PLUS:
		return Monkey{operation: MINUS, left: closed, right: m.left}
	case MINUS:
		return Monkey{operation: MINUS, left: m.left, right: closed}
	case DIV:
		return Monkey{operation: DIV, left: m.left, right: closed}
	case MUL:
		return Monkey{operation: DIV, left: closed, right: m.left}
	default:
		return Monkey{}
	}
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

func dependsOnHuman(name string, monkeys map[string]Monkey) bool {
	if name == HUMAN {
		return true
	}
	m := monkeys[name]
	if m.operation == "" {
		return false
	}
	return dependsOnHuman(m.left, monkeys) || dependsOnHuman(m.right, monkeys)
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
	var open, closed string
	if dependsOnHuman(monkeys[ROOT].left, monkeys) {
		open = monkeys[ROOT].left
		closed = monkeys[ROOT].right
	} else {
		open = monkeys[ROOT].right
		closed = monkeys[ROOT].left
	}
	for open != HUMAN {
		m := monkeys[open]
		if dependsOnHuman(m.left, monkeys) {
			uuid, _ := uuid.NewRandom()
			newname := uuid.String()
			monkeys[newname] = m.lhs(closed)
			closed = newname
			open = m.left
		} else {
			uuid, _ := uuid.NewRandom()
			newname := uuid.String()
			monkeys[newname] = m.rhs(closed)
			closed = newname
			open = m.right
		}
	}
	result := solve(monkeys[closed], monkeys)
	fmt.Println(result)
}
