package main

import "fmt"

func main() {
	var N, K int

	fmt.Scan(&N)
	fmt.Scan(&K)

	fmt.Println(N*(N-1)/2 - K)
}
