/* Exercise: Slices */
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	d2arr := make([][]uint8,dy)
	
	for i := 0; i < dy; i++ {
		d2arr[i] = make([]uint8,dx)
		for j := 0; j < dx; j++ {
			d2arr[i][j] = uint8((i + j)/2)
		}
	}
	return d2arr
}

func main() {
	pic.Show(Pic)
}