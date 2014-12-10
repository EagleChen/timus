package main

import "fmt"

func main() {
	var N, M uint64
	fmt.Scan(&N)
	fmt.Scan(&M)

	result1 := 2*N - 2
	result2 := 2*M - 1

	if result1 < result2 {
		fmt.Println(result1)
	} else {
		fmt.Println(result2)
	}
}
