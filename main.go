// Copyright (c) 2021 Quirino Gervacio

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// system to use
	sys := 6

	// how many set to generate
	set := 24

	// any numbers to skip
	skip := []int{}

	for i := 1; i <= set; i++ {
		var bet []int

		// don't stop until I get set of sys
		for {
			// toto range
			n := random(1, 49)

			// check for skips
			if exist(skip, n) {
				// fmt.Printf("Skipping %d\n", n)
				continue
			}

			// check if n is already generated
			if exist(bet, n) {
				// fmt.Printf("Already exist %d\n", n)
				continue
			}

			bet = append(bet, n)
			// fmt.Printf("Added %d\n", n)

			// if set is full
			if len(bet) == sys {
				sort.Ints(bet)
				fmt.Printf("Set#%d done. Got %v\n", i, bet)
				break
			}
		}
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

func exist(slice []int, n int) bool {
	for _, i := range slice {
		if i == n {
			return true
		}
	}
	return false
}
