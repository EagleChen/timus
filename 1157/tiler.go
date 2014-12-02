package main

import (
	"fmt"
	"math"
)

func getCount(num int) int {
	max := int(math.Floor(math.Sqrt(float64(num))))
	count := 0
	for i := 1; i <= max; i++ {
		if num%i == 0 {
			count += 1
		}
	}
	return count
}

func main() {
	var M, N, K int

	fmt.Scan(&M)
	fmt.Scan(&N)
	fmt.Scan(&K)

	counts := [10001]int{}
	for i := 1; i <= 10000; i++ {
		counts[i] = getCount(i)
	}

	for i, j := 1, 1+K; j <= 10000; i, j = i+1, j+1 {
		if counts[i] == M && counts[j] == N {
			fmt.Println(j)
			return
		}
	}

	fmt.Println(0)
}
