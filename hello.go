package main

import (
	"fmt"
      	//"newmath"
	"math/cmplx"
	"math"
)
var i, j int = 1, 2
var c, python, java = true, false, "no!"
var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)
func add(x, y int) int {
	return x + y
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main(){
// use := without var dec
	k := 3
//	const f = "%T(%v)\n"
	fmt.Println("Hello World! Sqrt(2) = %v\n",math.Sqrt(2))
	fmt.Println(add(42,k))
	fmt.Println(split(17))
	fmt.Println(i, j, c, python, java)
//	fmt.Printf(f, ToBe, ToBe)
//	fmt.Printf(f, MaxInt, MaxInt)
//	fmt.Printf(f, z, z)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, z)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	
	fmt.Println(sum)
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
