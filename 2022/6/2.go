package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, _ := ioutil.ReadFile("input.txt")
	bfr := string(b)
	offset := 14
	for i := range bfr {
		views := make(map[rune]bool)
		if i >= offset {
			for _, c := range bfr[i-offset : i] {
				if _, ok := views[c]; ok {
					break
				}
				views[c] = true
			}
		}
		if len(views) == offset {
			fmt.Println(i)
			return
		}
	}
}
