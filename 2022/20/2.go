package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	FST int = 1000
	SCD int = 2000
	THD int = 3000
	KEY int = 811589153
)

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	size int
}

func mixer(list *LinkedList, moveOrder []*Node) {
	for _, current := range moveOrder {
		offset := current.data % (list.size - 1)
		if offset == 0 {
			continue
		} else if offset < 0 {
			current.prev.next = current.next
			current.next.prev = current.prev
			insert := current
			for dx := 0; dx < -offset; dx++ {
				insert = insert.prev
			}
			insert.prev.next = current
			current.prev = insert.prev
			current.next = insert
			insert.prev = current
		} else {
			current.prev.next = current.next
			current.next.prev = current.prev
			insert := current
			for dx := 0; dx < offset; dx++ {
				insert = insert.next
			}
			insert.next.prev = current
			current.next = insert.next
			insert.next = current
			current.prev = insert
		}
	}
}

func main() {
	list := &LinkedList{}
	var current *Node
	var result int
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		newCurrent := &Node{data: n, prev: current}
		if list.head == nil {
			list.head = newCurrent
		} else {
			current.next = newCurrent
		}
		current = newCurrent
		list.size++
	}
	current.next = list.head
	list.head.prev = current
	for i := 0; i < list.size; i++ {
		list.head.data *= KEY
		list.head = list.head.next
	}
	moveOrder := []*Node{}
	for i := 0; i < list.size; i++ {
		moveOrder = append(moveOrder, list.head)
		list.head = list.head.next
	}
	for i := 0; i < 10; i++ {
		mixer(list, moveOrder)
	}
	for list.head.data != 0 {
		list.head = list.head.next
	}
	for i := 0; i <= THD; i++ {
		if i == FST || i == SCD || i == THD {
			result += list.head.data
		}
		list.head = list.head.next
	}
	fmt.Println(result)
}
