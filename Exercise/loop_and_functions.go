package main

import (
	"fmt"
	"math"
)

/* Exercise: Loops and Functions */
func Sqrt10(x float64) float64 {
	z := float64(1)
	for i:=0; i<10; i++ {
		z -= (z*z -x) / (2*z)
	}
	return z
}

func Sqrt(x float64) float64 {
	z := float64(1)
	prez := float64(0)
	counter := 0
	for  math.Abs(z - prez) > 1e-10 {
		prez = z
		z -= (z*z -x) / (2*z)
		counter += 1
	}
	fmt.Printf("loop count : %d, value : %g \n", counter, z)

	return z
}