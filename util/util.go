package util

import "math"

func Pow(b int, p int) int {
	return int(math.Pow(float64(b), float64(p))) - 1
}

func Abs(num int) int {
	return int(math.Abs(float64(num)))
}
