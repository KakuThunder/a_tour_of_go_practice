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

/* function */
func Absf(v Point) float64 {
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

	//func　の方はポインタを引数に取る
	v2 := Point{3, 4}
	v2.Scale(2)
	Scalef(&v2, 10)

	p2 := &Point{4, 3}
	p2.Scale(3)
	Scalef(p2, 8)

	fmt.Println(v2, p2)

	v3 := Point{3, 4}
	fmt.Println(v3.Absm())
	fmt.Println(Absf(v3))

	p3 := &Point{4, 3}
	fmt.Println(p3.Absm())
	fmt.Println(Absf(*p3))

}