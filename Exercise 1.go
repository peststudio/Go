package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1.0
    
    for i := 0; i < 1000; i++ {
        z = z - (math.Pow(z,2) - x) / (z * 2)
        if z == math.Sqrt(x) {
            fmt.Println("Number of Iterations ", i)
            return z
        }
    }
    return z
}

func main() {
    fmt.Println("My Function returns ", Sqrt(17))
    fmt.Println("math.Sqrt() returns ", math.Sqrt(17))
}
