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
type Pair struct {
    X int
    Y int
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
	if sum == 1 {
	fmt.Println("sum is 1")
	} else {
	fmt.Println("sum is not 1")
	}

	//struct
	v := Pair{1, 2}
    	v.X = 4
   	 v.Y = 5
    	fmt.Println(Pair(v))

	p := Pair{1, 2}
   	q := &p
    	q.X = 1e9
    	fmt.Println(p)
	
    	p = Pair{1, 2}  // has type Vertex
   	q = &Pair{1, 2} // has type *Vertex
   	p = Pair{X: 1}  // Y:0 is implicit
    	p = Pair{}      // X:0 and Y:0
	
	var a [2]string
    	a[0] = "Hello"
    	a[1] = "World"
    	fmt.Println(a[0], a[1])
    	fmt.Println(a)

	r := []int{2, 3, 5, 7, 11, 13}
    	fmt.Println("r ==", r)

    	for i := 0; i < len(r); i++ {
        	fmt.Printf("r[%d] == %d\n", i, r[i])
    	}
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	b = b[:cap(b)] // len(b)=5, cap(b)=5
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	type Vertex struct {
   		Lat, Long float64
	}

	var m = map[string]Vertex{
    		"Bell Labs": {40.68433, -74.39967},
    		"Google":    {37.42202, -122.08408},
	}
	func main() {
    		for i, v := range pow {
       		fmt.Printf("2**%d = %d\n", i, v)
		fmt.Println(m)
    	}
}
