package main

import (
	"fmt"
)

func main() {
	flower := []int{1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1}
	planetFlower(flower, 5)
}

func planetFlower(flowered []int, n int) {
	c := 0
	for i := 0; i < len(flowered); i++ {
		if flowered[i] == 0 {
			prevFlower := (i == 0 || flowered[i-1] == 0)
			nextFlower := (i == len(flowered)-1 || flowered[i+1] == 0)

			if prevFlower && nextFlower {
				c++
				i++
			}
		}
	}

	if n > c {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}
}
