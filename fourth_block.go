package main

/*
Methods and interfaces
*/

import (
	"fmt"
	"io"
	"math"
	"strings"
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

//Interface

type Inter interface {
	Met()
}
type Typ struct {
	Strin string
}

func (t *Typ) Met() {
	if t == nil{
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.Strin)
}

type Flo float64

func (f Flo) Met() {
	fmt.Println(f)
}

func describe(i Inter) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func empdescribe(i interface{}){
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v , v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!", v)
	}
}

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

//----------- main ----------------
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

	//変数はポインタをとれない
	v3 := Point{3, 4}
	fmt.Println(v3.Absm())
	fmt.Println(Absf(v3))

	p3 := &Point{4, 3}
	fmt.Println(p3.Absm())
	fmt.Println(Absf(*p3))

	v4 := &Point{3, 4}
	fmt.Printf("Before scaling: %+v, Absm: %v\n", v4, v4.Absm())
	v4.Scale(5)
	fmt.Printf("After scaling: %+v, Absm: %v\n", v4, v4.Absm())

	var interf Inter = &Typ{"Hello"}
	interf.Met()

	interf = Flo(math.Pi)
	describe(interf)
	interf.Met()

	//methodがnilをレシーバとしても呼び出せる
	var tmp *Typ
	interf = tmp
	describe(interf)
	interf.Met()

	//nilインタフェースのメソッドを呼ぶことはできない
	/*
	var i I
	describe(i)
	i.M()
	*/

	//メソッドが指定されていないものは空interface
	//任意の型の値を保持することができる
	//型が宣言時に判断できないものを取り扱うときに利用する
	var empi interface{}
	empdescribe(empi)
	empi = 42
	empdescribe(empi)
	empi = "hello"
	empdescribe(empi)

	//型アサーション
	var assi interface{} = "String"

	s := assi.(string)
	fmt.Println(s)

	s,ok := assi.(string)
	fmt.Println(s, ok)

	fl,ok := assi.(float64)
	fmt.Println(fl,ok)

	//fl = assi.(float64) //panic
	//fmt.Println(fl)

	//type swirch
	//様々な型のデータをとれる
	do(21)
	do("Hello")
	do(true)

	fmt.Println()

	humana := Person{"Arthur Dent", 42}
	humanz := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(humana,humanz)

	comment := strings.NewReader("Hello, Reader!")

	block := make([]byte,8)
	for {
		n, err := comment.Read(block)
		fmt.Printf("n=%v err = %v b = %v\n", n, err, block)
		fmt.Printf("block[:n] = %q\n",block[:n])

		if err == io.EOF {
			break
		}
	}


}