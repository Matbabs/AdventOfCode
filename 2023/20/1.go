package main

import (
	"bufio"
	"fmt"
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

func main() {
	mlow, mhigh := 0, 0
	msgs := []Msg{}
	var brd Pulse
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
		}
	}
	for b := 0; b < 1000; b++ {
		mlow++
		for _, e := range brd.exits {
			msgs = append(msgs, Msg{brd.name, 0, e})
		}
		for len(msgs) > 0 {
			m := msgs[0]
			msgs = msgs[1:]
			if m.v == 0 {
				mlow++
			} else {
				mhigh++
			}
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
	fmt.Println(mlow * mhigh)
}
