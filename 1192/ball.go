package main

import (
	"fmt"
	"math"
)

func main() {
	var v, a, k, distance, g float64

	g = 10
	p := 3.1415926535

	fmt.Scan(&v)
	fmt.Scan(&a)
	fmt.Scan(&k)

	a = p * a / 180

	for {
		if v < 0.01 {
			break
		}
		distance += v * v * math.Sin(2*a) / g
		v /= math.Sqrt(k)
	}

	fmt.Printf("%.2f\n", distance)
}
