package main

/*
More types:
	structs
	slices
	maps
*/

import (
	"fmt"
)

//Struct
type Vertex struct {
    X int
    Y int
}


//構造体の初期化
var (
    v1 = Vertex{1,2}
    v2 = Vertex{X: 1}
    v3 = Vertex{}
    p3 = &Vertex{1,2}
)

func main3() {

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

}