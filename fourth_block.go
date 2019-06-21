package main

/*
Methods and interfaces
*/

import (
	"fmt"
	"math"
)


type Point struct {
	X, Y float64
}
/* Methods */
func (v Point) Absm() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
/* Pointer receivers */
//pointerでないと関数内のみのスコープになる
func (v *Point) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f	
}


func Scalef(v *Point, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f	
}

/* function */
func Absf(v Point) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/* Methods continued */
type MyFloat float64

func (f MyFloat) Absmc() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}


func main4(){

	fmt.Println("------ start main4 ------")

	v := Point{3,4}
	fmt.Println(v.Absm())
	fmt.Println(Absf(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Absmc())

	v.Scale(10)
	fmt.Println(v.Absm())

	Scalef(&v,10)
	fmt.Println(Absf(v))

}