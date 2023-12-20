package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

type Pulse struct {
	name    string
	on      bool
	entries []string
	exits   []string
	memory  map[string]int
	def     string
}

type Msg struct {
	entry string
	v     int
	exit  string
}

func lcm(buttons []int64) *big.Int {
	lcm := big.NewInt(buttons[0])
	for i := 1; i < len(buttons); i++ {
		gcd := new(big.Int)
		gcd.GCD(nil, nil, lcm, big.NewInt(buttons[i]))
		lcm.Mul(lcm, big.NewInt(buttons[i]))
		lcm.Div(lcm, gcd)
	}
	return lcm
}

func main() {
	btn := 0
	msgs := []Msg{}
	var brd Pulse
	rxEntries := []Pulse{}
	rxAcq := make(map[string]int)
	pls := []Pulse{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ls := strings.Split(scanner.Text(), " ")
		nm, d := ls[0], ""
		if strings.Contains(nm, "%") {
			nm = strings.ReplaceAll(nm, "%", "")
			d = "%"
		}
		if strings.Contains(nm, "&") {
			nm = strings.ReplaceAll(nm, "&", "")
			d = "&"
		}
		ex := []string{}
		for i := 2; i < len(ls); i++ {
			ex = append(ex, strings.ReplaceAll(ls[i], ",", ""))
		}
		pls = append(pls, Pulse{name: nm, on: false, exits: ex, def: d, memory: make(map[string]int)})
	}
	for i := range pls {
		if pls[i].name == "broadcaster" {
			brd = pls[i]
		}
		for _, e := range pls[i].exits {
			for j := range pls {
				if pls[j].name == e {
					pls[j].entries = append(pls[j].entries, pls[i].name)
					if pls[j].def == "&" {
						pls[j].memory[pls[i].name] = 0
					}
				}
			}
			if e == "rx" {
				rxEntries = append(rxEntries, pls[i])
			}
		}
	}
	for {
		btn++
		for _, e := range brd.exits {
			msgs = append(msgs, Msg{brd.name, 0, e})
		}
		for len(msgs) > 0 {
			nbAcq := 0
			for i := range rxEntries {
				for rxE, v := range rxEntries[i].memory {
					if v == 1 {
						if _, p := rxAcq[rxE]; !p {
							rxAcq[rxE] = btn
						}
					}
				}
				nbAcq += len(rxEntries[i].memory)
			}
			if len(rxAcq) == nbAcq {
				acqBtn := []int64{}
				for _, b := range rxAcq {
					acqBtn = append(acqBtn, int64(b))
				}
				fmt.Println(lcm(acqBtn))
				return
			}
			m := msgs[0]
			msgs = msgs[1:]
			for i := range pls {
				if pls[i].name == m.exit {
					switch pls[i].def {
					case "%":
						if m.v == 0 {
							pls[i].on = !pls[i].on
							v := 0
							if pls[i].on {
								v = 1
							}
							for _, e := range pls[i].exits {
								msgs = append(msgs, Msg{pls[i].name, v, e})
							}
						}
					case "&":
						pls[i].memory[m.entry] = m.v
						v := 1
						hgh := true
						for _, m := range pls[i].memory {
							if m == 0 {
								hgh = false
								break
							}
						}
						if hgh {
							v = 0
						}
						for _, e := range pls[i].exits {
							msgs = append(msgs, Msg{pls[i].name, v, e})
						}
					}
				}
			}
		}
	}
}
