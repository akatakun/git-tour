package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// 関数内のみで利用可能な暗黙的な型宣言
	// var k int = 3
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
