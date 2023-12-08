package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"slices"
	"strings"
)

func lcm(distances []int64) *big.Int {
	lcm := big.NewInt(distances[0])
	for i := 1; i < len(distances); i++ {
		gcd := new(big.Int)
		gcd.GCD(nil, nil, lcm, big.NewInt(distances[i]))
		lcm.Mul(lcm, big.NewInt(distances[i]))
		lcm.Div(lcm, gcd)
	}
	return lcm
}

func main() {
	p, l := "", 0
	a := []string{}
	lg := []int64{}
	f := []int{}
	n := make(map[string][]string)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if l == 0 {
			p = scanner.Text()
			l++
		} else if scanner.Text() != "" {
			nd := strings.Split(strings.ReplaceAll(scanner.Text(), " ", ""), "=")
			pp := strings.Split(nd[1], ",")
			n[nd[0]] = []string{pp[0][1:], pp[1][:3]}
		}
	}
	for nd := range n {
		if string(nd[2]) == "A" {
			a = append(a, nd)
			lg = append(lg, 0)
		}
	}
	for len(f) != len(a) {
		for _, r := range p {
			for i := range a {
				a[i] = n[a[i]][strings.IndexRune("LR", r)]
				if !slices.Contains(f, i) {
					lg[i]++
				}
				if string(a[i][2]) == "Z" {
					f = append(f, i)
				}
			}
			if len(f) == len(a) {
				break
			}
		}
	}
	fmt.Println(lcm(lg))
}
