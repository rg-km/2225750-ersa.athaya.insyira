package main

import (
	"fmt"
)

// Birthday Cake Candles
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/birthday-cake-candles/problem

func birthdayCakeCandles(candles []int32) int32 {
	// TODO: answer here
	var max int32 = 0
	var score = make(map[int32]int32)
	for _, val := range candles {
		if max < val {
			max = val
			score[max] = 1
		} else if max == val {
			score[max]++
		}
		fmt.Println("Nilai max :", max)
		fmt.Println("Nilai score : ", score)
	}
	return score[max]
}

func main() {
	var arr = []int32{3, 2, 1, 3}
	fmt.Println(birthdayCakeCandles(arr))
}
