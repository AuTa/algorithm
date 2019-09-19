package main

import "math"

/*
 * @lc app=leetcode id=829 lang=golang
 *
 * [829] Consecutive Numbers Sum
 */
func consecutiveNumbersSum(N int) int {
	var c int
	dn := 2 * N
	var sum int
	for i := 1; i <= int(math.Sqrt(float64(dn))); i++ {
		sum += i
		if (N-sum)%i == 0 {
			c++
		}
	}
	return c
}

func consecutiveNumbersSum2(N int) int {
	var c int
	dn := 2 * N
	for i := 1; i < int(math.Sqrt(float64(dn))); i++ {
		if dn%i != 0 {
			continue
		}
		divisor := dn / i
		if (divisor+i-1)&1 != 0 {
			continue
		}
		if (divisor-i+1)&1 != 0 {
			continue
		}
		c++
	}
	return c
}
