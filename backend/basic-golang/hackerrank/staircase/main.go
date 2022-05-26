package main

import (
	"fmt"
	"strings"
)

// Staircase
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/staircase/problem

func staircase(n int32) {
	// TODO: answer here
	for i := int32(1); i <= n; i++ {
		fmt.Printf("%*s\n", n, strings.Repeat("#", int(i)))
	}

}

func main() {
	staircase(10)
}
