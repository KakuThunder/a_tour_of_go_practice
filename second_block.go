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
	"math"
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


}