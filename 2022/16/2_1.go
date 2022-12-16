package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const START = "AA"
const MAX_TIME = 30

type Node struct {
	name      string
	rate      int
	edgesDist map[string]int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func visit(nodeMap map[string]Node, currentNode string, visited []string, time int, maxTime int) (int, []string) {
	if time >= maxTime {
		return 0, visited
	}
	node := nodeMap[currentNode]
	score := node.rate * (maxTime - time)
	newVisited := append(visited, currentNode)
	var maxScore int
	var maxPath []string
	for neighbor := range node.edgesDist {
		if !contains(visited, neighbor) {
			childScore, childPath := visit(nodeMap, neighbor, newVisited, time+node.edgesDist[neighbor]+1, maxTime)
			if childScore > maxScore {
				maxScore = childScore
				maxPath = childPath
			}
		}
	}
	return score + maxScore, append([]string{currentNode}, maxPath...)
}

func main() {
	nodes := make([]Node, 0)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		name := strings.Split(strings.Split(line, "Valve")[1], " ")[1]
		rate_str := strings.Split(strings.Split(line, "rate=")[1], ";")[0]
		neighbors := strings.Split(strings.Split(line, "valve")[1], " ")[1:]
		rate, _ := strconv.Atoi(rate_str)
		neighbors_map := make(map[string]int)
		for n := range neighbors {
			neighbors[n] = strings.TrimSpace(neighbors[n])
			neighbors[n] = strings.ReplaceAll(neighbors[n], ",", "")
			neighbors_map[neighbors[n]] = 1
		}
		nodes = append(nodes, Node{name, rate, neighbors_map})
	}
	for _, node1 := range nodes {
		for _, node2 := range nodes {
			_, ok := node2.edgesDist[node1.name]
			if node1.name != node2.name && ok {
				currentDistance := node2.edgesDist[node1.name]
				for neighbor, distance := range node1.edgesDist {
					if neighbor == node2.name {
						continue
					}
					if _, ok := node2.edgesDist[neighbor]; !ok {
						node2.edgesDist[neighbor] = currentDistance + distance
					} else {
						node2.edgesDist[neighbor] = min(node2.edgesDist[neighbor], currentDistance+distance)
					}
				}
				if node1.rate == 0 {
					delete(node2.edgesDist, node1.name)
				}
			}
		}
	}
	nodeMap := make(map[string]Node)
	for _, node := range nodes {
		if node.name == START || node.rate != 0 {
			nodeMap[node.name] = node
		}
	}
	maxFlow, path := visit(nodeMap, START, nil, 0, MAX_TIME)
	for _, node := range nodeMap {
		for _, p := range path[1:] {
			delete(node.edgesDist, p)
		}
	}
	elephantFlow, path := visit(nodeMap, START, nil, 0, MAX_TIME)
	fmt.Println(maxFlow + elephantFlow)
}
