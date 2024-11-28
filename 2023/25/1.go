package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gonum/graph"
	"github.com/gonum/graph/simple"
)

func main() {
	g := simple.NewUndirectedGraph()

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) > 0 {
			source := graph.Node(g.NodeCount())
			g.AddNode(source)

			for _, target := range parts[1:] {
				targetNode := graph.Node(g.NodeCount())
				g.AddNode(targetNode)
				g.SetEdge(simple.Edge{F: source, T: targetNode})
			}
		}
	}

	cv, p := stoerWagner(g)
	fmt.Println(len(p[0]) * len(p[1]))
}

func stoerWagner(g graph.Undirected) (int, []graph.NodeSet) {
	nodes := graph.NodesOf(g.Nodes())
	cv, p := simple.StoerWagner(g, nodes)
	return cv, p
}
