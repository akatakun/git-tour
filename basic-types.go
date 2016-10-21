package main

/*
基本型
bool
string
int(int8, int16, int32, int64)
uint(uint8, uint16, uint32, uint64, uintptr)
byte(alias of unit8)
rune(alias of int32)
float32, float64
complex64, complex128

特別な理由がない限りは整数にはなるべくint型を使う
*/

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// %T: type
	// %v: value
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
