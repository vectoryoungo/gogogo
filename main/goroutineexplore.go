package main

import (
	"strconv"
)

func main() {
	a := 1
	b := 2
	r := max(a, b)
	println(`max: ` + strconv.Itoa(r))
	const PtrSize = 4 << (^uintptr(0) >> 63)
	println(PtrSize)
	const fft = ^uintptr(0)
	println(fft)
	const shitWithSixThree = (^uintptr(0) >> 63)
	println(shitWithSixThree)
}

func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}
