package main

/*
Packages
variables
functions
*/

//個別にimportしても処理は変わらない
import (
	"fmt"
	"math"
	"math/rand"
	"math/cmplx"
)

// x int , y int と同様,同じ型であれば以下のように受け取れる
func add(x,y int) int{
	return x + y
}

//戻り値が複数なら()
func swap(x,y string) (string,string) {
	return y , x
}

//戻り値の変数名,型を事前に宣言することで,return時に自動的に値が返される
func split(sum int) (x,y int) {
	x = sum * 4/9
	y = sum -x 
	return
}


//定数は char, string, bool, num のみ 型ではなく,時々で必要な型をとる
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}





func main() {
	//しっかり乱数を生成するにはseedが必要
	//time.Now().UnixNano()をSeedにするなど
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println("Now you have %g problems.", math.Sqrt(7))

	//大文字開始の名前は,外部から参照していることを表す
	//Println,Intn,Sqrt,Pi
	fmt.Println(math.Pi)

	fmt.Println(add(45,23))

	//変数宣言なし
	a,b := swap("hello","world")
	fmt.Println(a,b)

	fmt.Println(split(17))

	//型宣言
	var c, python, java bool
	var tmpnum int
	var tmpval int = 5
	fmt.Println(tmpnum, tmpval, c, python, java)

	//型宣言を省略した初期化
	var today, yeesterday = 615,614
	var php, ruby, perl = true, false, "script"
	fmt.Println(today, yeesterday, php, ruby, perl)

	//関数内であればvarがなくとも行ける
	testval := 5
	ie,chrome,firefox := false,true,"hikistune"
	fmt.Println(testval, ie, chrome, firefox)

	//同型
	var minival byte = 255
	var minival2 uint8 = 255
	fmt.Printf("Type: %T Value %v\n", minival, minival)
	fmt.Printf("Type: %T Value %v\n", minival2, minival2)

	//同型
	var intval rune  = 1234
	var intval2 int32 = 4567
	fmt.Printf("Type: %T Value %v\n", intval, intval)
	fmt.Printf("Type: %T Value %v\n", intval2, intval)

	var kyosu complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Printf("Type: %T Value %v\n", kyosu, kyosu)

	var initi int
	var initf float64
	var initb bool
	var inits string
	fmt.Printf("%v %v %v %q\n", initi, initf, initb, inits)

	//型変換
	var changefloat = float32(intval)
	changedouble := float64(intval)
	fmt.Printf("Type: %T Value %v\n", changefloat, changefloat)
	fmt.Printf("Type: %T Value %v\n", changedouble, changedouble)

	//型推論
	whatval := 1024
	fmt.Printf("Type: %T Value %v\n", whatval, whatval)

	//定数宣言 (:= は利用不可)
	const effort = 512
	const comment = "努力値"

	fmt.Printf("%s の最大は %d\n", comment, effort)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))


}
