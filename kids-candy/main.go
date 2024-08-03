package main

import (
	"fmt"
	"sort"
)

func GreatestCandy(candy []int) int {
	sort.Ints(candy)
	return candy[len(candy)-1]
}

func sumCandies(candyarray []int, sum int) []int {
	var newCandy []int
	var extra int
	for i := 0; i < len(candyarray); i++ {
		extra = candyarray[i] + sum
		newCandy = append(newCandy, extra)
	}
	return newCandy
}

func CheckCandy(newCandy, oldCandy []int) []bool {
	var isTOrF []bool
	greatCandy := GreatestCandy(oldCandy)
	for i := 0; i < len(newCandy); i++ {
		if newCandy[i] >= greatCandy {
			isTOrF = append(isTOrF, true)
		} else {
			isTOrF = append(isTOrF, false)
		}
	}
	return isTOrF
}

func main() {
	candies := []int{1, 3, 5, 9, 4, 7}
	extraCandies := 4
	sumCandy := sumCandies(candies, extraCandies)
	checkIt := CheckCandy(sumCandy, candies)
	fmt.Println(checkIt)
}
