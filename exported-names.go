package main

// パッケージがインポートされるとエクスポートされた名前(大文字から始まる)を直接参照できるようになる
// エクスポートされていない名前(小文字から始まる)はプライベートみないなもの
import (
	"fmt"
	"math"
)

/*
# command-line-arguments
./exported-names.go:11: cannot refer to unexported name math.pi
./exported-names.go:11: undefined: math.pi
func main() {
	fmt.Println(math.pi)
}
*/

func main() {
	fmt.Println(math.Pi)
}
