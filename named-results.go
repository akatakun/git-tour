package main

import "fmt"

// returnでxとyが自動的に返される(naked return)
// 短い関数でのみ使うべき
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
