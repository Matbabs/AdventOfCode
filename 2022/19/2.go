package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
)

var DP map[DPKey]int
var hits = 0

type Plan struct {
	Blueprints []Blueprint
}

type ObsCost struct {
	OreCost  int
	ClayCost int
}

func (o ObsCost) String() string {
	return fmt.Sprintf("o=%d, c=%d", o.OreCost, o.ClayCost)
}

type GeodeCost struct {
	OreCost int
	ObsCost int
}

func (g GeodeCost) String() string {
	return fmt.Sprintf("o=%d, ob=%d", g.OreCost, g.ObsCost)
}

type Blueprint struct {
	OreCost   int
	ClayCost  int
	ObsCost   ObsCost
	GeodeCost GeodeCost
}

func (b Blueprint) String() string {
	return fmt.Sprintf("o=%d c=%d ob={%v}, g={%v}", b.OreCost, b.ClayCost, b.ObsCost, b.GeodeCost)
}

type DPKey struct {
	t, ores, clays, obsidians, oreRobots, clayRobots, obsidianRobots, geodeRobots int
}

func (d DPKey) String() string {
	return fmt.Sprintf("t=%d, quants=[%d, %d, %d], robots=[%d, %d, %d, %d]", d.t, d.ores, d.clays, d.obsidians, d.oreRobots, d.clayRobots, d.obsidianRobots, d.geodeRobots)
}

func allInts(s string) []int {
	res := make([]int, 0)
	re := regexp.MustCompile(`\d+`)
	for _, n := range re.FindAllString(s, -1) {
		v, _ := strconv.Atoi(n)
		res = append(res, v)
	}
	return res
}

func max(ints ...int) int {
	max := 0
	for _, i := range ints {
		if i > max {
			max = i
		}
	}
	return max
}

func geodeCount(b Blueprint, k DPKey) int {
	if val, ok := DP[k]; ok {
		hits++
		return val
	}
	if k.t == 0 {
		return 0
	}
	best := k.geodeRobots
	if k.ores >= b.GeodeCost.OreCost && k.obsidians >= b.GeodeCost.ObsCost {
		best = max(best, k.geodeRobots+geodeCount(b, DPKey{
			k.t - 1,
			k.ores + k.oreRobots - b.GeodeCost.OreCost, k.clays + k.clayRobots, k.obsidians + k.obsidianRobots - b.GeodeCost.ObsCost,
			k.oreRobots, k.clayRobots, k.obsidianRobots, k.geodeRobots + 1,
		}))
	} else {
		if k.ores >= b.ObsCost.OreCost && k.clays >= b.ObsCost.ClayCost {
			best = max(best, k.geodeRobots+geodeCount(b, DPKey{
				k.t - 1,
				k.ores + k.oreRobots - b.ObsCost.OreCost, k.clays + k.clayRobots - b.ObsCost.ClayCost, k.obsidians + k.obsidianRobots,
				k.oreRobots, k.clayRobots, k.obsidianRobots + 1, k.geodeRobots,
			}))
		}
		if k.ores >= b.ClayCost {
			best = max(best, k.geodeRobots+geodeCount(b, DPKey{
				k.t - 1,
				k.ores + k.oreRobots - b.ClayCost, k.clays + k.clayRobots, k.obsidians + k.obsidianRobots,
				k.oreRobots, k.clayRobots + 1, k.obsidianRobots, k.geodeRobots,
			}))
		}
		if k.ores >= b.OreCost {
			best = max(best, k.geodeRobots+geodeCount(b, DPKey{
				k.t - 1,
				k.ores + k.oreRobots - b.OreCost, k.clays + k.clayRobots, k.obsidians + k.obsidianRobots,
				k.oreRobots + 1, k.clayRobots, k.obsidianRobots, k.geodeRobots,
			}))
		}

		maxOres := max(b.OreCost, b.ClayCost, b.ObsCost.OreCost, b.GeodeCost.OreCost)
		if k.ores < maxOres || k.clays < b.ClayCost || k.obsidians < b.GeodeCost.ObsCost {
			best = max(best, k.geodeRobots+geodeCount(b, DPKey{
				k.t - 1,
				k.ores + k.oreRobots, k.clays + k.clayRobots, k.obsidians + k.obsidianRobots,
				k.oreRobots, k.clayRobots, k.obsidianRobots, k.geodeRobots,
			}))
		}
	}
	DP[k] = best
	return best
}

func main() {
	result := 0
	plan := &Plan{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x := allInts(scanner.Text())
		plan.Blueprints = append(plan.Blueprints, Blueprint{
			OreCost:  x[1],
			ClayCost: x[2],
			ObsCost: ObsCost{
				OreCost:  x[3],
				ClayCost: x[4],
			},
			GeodeCost: GeodeCost{
				OreCost: x[5],
				ObsCost: x[6],
			},
		})
	}
	for i, b := range plan.Blueprints {
		if i == 3 {
			break
		}
		DP = map[DPKey]int{}
		runtime.GC()
		result *= geodeCount(b, DPKey{32, 0, 0, 0, 1, 0, 0, 0})
	}
	fmt.Println(result)
}
