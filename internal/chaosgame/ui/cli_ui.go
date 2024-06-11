package ui

import "fmt"

func PrintCanvas(canvas [][]bool) {
	for _, row := range canvas {
		for _, val := range row {
			if val {
				fmt.Printf("■")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
