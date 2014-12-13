package main

import "fmt"

func main() {
	var K int
	fmt.Scan(&K)

	var noLeading uint64 = 55
	var leading uint64 = 36
	var result uint64 = leading
	for i := K - 1; i > 0; i-- {
		result *= noLeading
	}
	fmt.Println(result)
}
