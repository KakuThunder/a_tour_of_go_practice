package main

import "fmt"



func main(){

	/* Exercise: Loops and Functions */	
	arg := 1
	for; arg < 10; {
		fmt.Println(Sqrt10(float64(arg)))
		arg +=1
	}

	counter := 1
	for; counter < 10; {
		fmt.Println(Sqrt(float64(counter)))
		counter +=1
	}


}