package main

import "fmt"
import "sort"

func main() {
	var N int
	fmt.Scan(&N)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Ints(arr)

	for i := N - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}
}
