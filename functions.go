package main

import (
	"C"
	"math"
)

//export IsSimple
func IsSimple(n int) bool {
	isSimple := true
	if (n == 0) || (n == 1) {
		return false
	}
	for i := 2; int64(i) <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			isSimple = false
		}
	}
	return isSimple
}

//export NextSimple
func NextSimple(i int) int {
	isFind := false
	digit := i + 1
	for !isFind {
		if IsSimple(digit) {
			return digit
		}
		digit++
	}
	return 0
}
func main() {}
