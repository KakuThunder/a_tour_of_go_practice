package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fi1,fi2 := 0,1
	return func () int {
		retval := fi1
		fi1,fi2 = fi2,fi1+fi2
		return retval
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}