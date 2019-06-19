package main

/*
More types:
	structs
	slices
	maps
*/

import (
	"fmt"
	"math"
	"strings"
)

//Struct
type Vertex struct {
    X int
    Y int
}
type Location struct {
	Lat, Long float64
}



//構造体の初期化
var (
    v1 = Vertex{1, 2}
    v2 = Vertex{X: 1}
    v3 = Vertex{}
    p3 = &Vertex{1, 2}
)

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printmakeSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

//関数を引数とする関数
func compute(fn func(float64, float64) float64) float64 {
	return fn(3,4)
}

func adder() func(x int) int{
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
 
func main3() {

	fmt.Println("------ start main3 ------")

	//Pointers
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p /37
	fmt.Println(j)

	//構造体の宣言, 構造体のフィールドへのアクセス
	fmt.Println(Vertex{1,2})
	v := Vertex{1,2}
	v.X = 4
	fmt.Println(v.X)

	//ポインタでの構造体フィールドへのアクセス
	ve := Vertex{1,2}
	p2 := &ve
	p2.X = 1e9
	fmt.Println(ve)

	fmt.Println(v1,v2,v3,p3)

	//Array
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0],a[1])
	fmt.Println(a)

	primes := [6]int {2,3,5,7,11,13}
	fmt.Println(primes)

	//配列のスライス
	var s []int = primes[1:4]
	fmt.Println(s)

	//スライスは配列の参照であり
	//新たに生成しているわけではない
	names := [4]string {
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	aslice := names[0:2]
	bslice := names[1:3]
	fmt.Println(aslice,bslice)

	bslice[0] = "XXX"
	fmt.Println(aslice,bslice)
	fmt.Println(names)	

	//スライスのリテラル
	//配列を生成して参照するスライスを作成する
	qarr := []int {2,3,5,7,11,13}
	fmt.Println(qarr)

	r := []bool {true, false, true, true, false, true}
	fmt.Println(r)

	strux := []struct {
		inum int
		flag bool
	} {
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(strux)

    slarr := []int{2, 3, 5, 7, 11, 13}

    slarr = slarr[1:4]
    fmt.Println(slarr)

    slarr = slarr[:2]
    fmt.Println(slarr)

    slarr = slarr[1:]
    fmt.Println(slarr)

    /*Slice length and capacity*/
    funsarr := []int{2, 3, 5, 7, 11, 13}
    printSlice(funsarr)

    // Slice the slice to give it zero length.
    funsarr = funsarr[:0]
    printSlice(funsarr)

    // Extend its length.
    funsarr = funsarr[:4]
    printSlice(funsarr)

    // Drop its first two values.
    funsarr = funsarr[2:]
	printSlice(funsarr)
	
	//空スライス nil
	var nils []int
	fmt.Println(s, len(s), cap(s))
	if nils == nil {
		fmt.Println("nil")
	}

	//動的配列作成
	//型 , 長さ, 容量
	slicea := make([]int, 5)
	printmakeSlice("slicea", slicea)

	sliceb := make([]int, 0, 5)
	printmakeSlice("sliceb", sliceb)

	slicec := sliceb[:2]
	printmakeSlice("slicec", slicec)

	sliced := slicec[2:5]
	printmakeSlice("sliced", sliced)

	board := [][]string {
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[1][1] = "X"
	board[0][1] = "O"
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][0] = "X"
	board[1][2] = "O"
	board[2][0] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	//slice への要素の追加
	var appends []int
	printSlice(appends)

	appends = append(appends, 0)
	printSlice(appends)

	appends = append(appends, 1)
	printSlice(appends)

	appends = append(appends, 2, 3, 4)
	printSlice(appends)

	//slice の for文処理
	var itrpow = []int {1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}

	// index , value := range(slice/map)
	for itr1, itr2 :=  range itrpow {
		fmt.Printf("2**%d = %d \n", itr1, itr2)
	}

	rangepow := make([]int, 10)
	for i := range rangepow {
		rangepow[i] = 1 << uint(i) 
	}
	for _, value := range rangepow {
		fmt.Printf("%d\n", value)
	}

	//map 
	var point map[string] Location

	point = make(map[string]Location)
	point["Bell Labs"] = Location{
		40.68433, -74.39967,
	}
	fmt.Println(point["Bell Labs"])

	/*
	var point2 = map[string] Location {
		"Bell Labs": Location {
			40.68433, -74.39967,
		},
		"Google": Location {
			37.42202, -122.08408,
		},		
	}*/
	var point2 = map[string] Location{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(point2)

	mutatem := make(map[string]int)

	mutatem["Answer"] = 42
	fmt.Println("The value:", mutatem["Answer"])

	mutatem["Answer"] = 48
	fmt.Println("The value:", mutatem["Answer"])

	delete(mutatem, "Answer")
	fmt.Println("The value:", mutatem["Answer"])

	value, ok := mutatem["Answer"]
	fmt.Println("The value:", value, "Present?", ok)	

	hypot := func(x,y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5,12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++{
		fmt.Println(pos(i),neg(-2*i))
	}
}