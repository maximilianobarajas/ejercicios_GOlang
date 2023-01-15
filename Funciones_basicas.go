
package main

import (
	"fmt"
	"math"
)
func correr(fn func(float64, float64) float64, h,g float64) float64 {
	return fn(h, g)
}

func hypotenusa(x, y float64)float64 {
		return math.Sqrt(x*x + y*y)
	}
func main() {
	fmt.Println("El valor de la hipotenusa es: ",hypotenusa(5, 12))
	fmt.Println("El valor de la hipotenusa es: ",correr(hypotenusa,5,12))
    fmt.Println("El valor de la potenciación es: ",correr(math.Pow,5,12))
}
