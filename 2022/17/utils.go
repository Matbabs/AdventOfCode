package main

import "fmt"

func draw(sx, sy, actual_shape, max_height int, tower map[string]bool, isMoving bool) {
	for i := max_height + sy; i >= 0; i-- {
		for j := 0; j < WIDTH; j++ {
			_, ok := tower[fmt.Sprintf("%d,%d", j, i)]
			if ok {
				fmt.Print("#")
			} else {
				if isMoving {
					draw := false
					for _, block := range shapes[actual_shape] {
						px, py := sx+block[0], sy+block[1]
						if px == j && py == i {
							fmt.Print("@")
							draw = true
						}
					}
					if !draw {
						fmt.Print(".")
					}
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}
