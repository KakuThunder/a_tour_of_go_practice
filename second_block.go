package main

/*
Flow control statements:
	for
	if
	else
	switch
	defer
*/

import (
	"fmt"
	"time"
	"math"
	"runtime"
)

//if statement pattern1
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))

}

//if statement pattern2
func pow(x,n,lim float64) float64 {
	if v := math.Pow(x,n); v < lim {
		return v
	}
	return lim
}

//if else 
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}


func main2() {

	fmt.Println("------ start main2 ------")

	//defer に設定した関数の実行を関数の終了の直前に行ってくれる
	//関数の引数は設定した時点での値を評価することを留意しておく
	//複数ある場合は <LIFO> で実行
	defer fmt.Println("proc end")
	defer fmt.Println("process complete")

	//for roop pattern1
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//for roop pattern2
	nums:=1
	for; nums < 1000; {
		nums += nums
	}
	fmt.Println(nums)

	//for roop pattern3 
	num2 := 1
	for num2 < 1000 {
		num2 += num2
	}
	fmt.Println(num2)

	//無限ループ
	//for{}

	fmt.Println(sqrt(2),sqrt(-4))	

	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
	)

	fmt.Println(
		pow2(3,2,10),
		pow2(3,3,20),
	)

	//switch 一番最初に条件を満たすcaseのみを実行する
	switch os:= runtime.GOOS; os{
		case "darwin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux.")
		default:
			//freebsd, openbsd,
			//plan9, windows...
			fmt.Printf("%s.\n",os)
	}

	today := time.Now().Weekday()
	switch time.Saturday {
		case today + 0:
			fmt.Println("Today.")
		case today + 1:
			fmt.Println("Tommorow.")
		case today + 2:
			fmt.Println("In too days.")
		default:
			fmt.Println("Too far away.")
	}

	//一番最初のcaseのみ実行されるので
	//XX以上,XX以下が簡単に記載できる
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("Good morning!")
		case t.Hour() < 17:
			fmt.Println("Good afternoon!")
		default:
			fmt.Println("Good evening!")
	}


}