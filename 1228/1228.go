package main

import "fmt"

func main() {
	var n, total int
	fmt.Scan(&n)
	fmt.Scan(&total)

	multipliers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&multipliers[i])
	}

	boundaries := make([]int, n)
	for i := n - 1; i > 0; i-- {
		boundaries[i] = multipliers[i-1]/multipliers[i] - 1
	}
	boundaries[0] = total/multipliers[0] - 1

	output := fmt.Sprint(boundaries)
	fmt.Println(output[1 : len(output)-1])
}
