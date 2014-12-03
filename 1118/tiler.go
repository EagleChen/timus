package main

import (
	"fmt"
	"math"
)

func triviality(num int) (float64, bool) {
	sum := 1
	max := int(math.Floor(math.Sqrt(float64(num))))
	for i := 2; i <= max; i++ {
		other := num % i
		if other == 0 {
			sum += i + num/i
		}
	}
	isPrime := false
	if sum == 1 {
		isPrime = true
	}
	return float64(sum) / float64(num), isPrime
}

func main() {
	var I, J int

	fmt.Scan(&I)
	fmt.Scan(&J)

	if I == 1 {
		fmt.Println(1)
		return
	}
	min := float64(1000000)
	result := 1000001
	for i := J; i >= I; i-- {
		t, isPrime := triviality(i)
		if isPrime {
			result = i
			break
		}
		if t < min {
			min = t
			result = i
		}
	}

	fmt.Println(result)
}
