package main

import "fmt"
import "math"
import "strings"

func value(index int) string {
	if index == 1 {
		return "1"
	}

	diff := 2 * (index - 1)
	root := int(math.Floor(math.Sqrt(float64(diff))))
	if diff%root == 0 && diff/root-root == 1 {
		return "1"
	} else {
		return "0"
	}
}

func main() {
	var N, index int
	fmt.Scan(&N)

	result := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&index)
		result[i] = value(index)
	}
	fmt.Println(strings.Join(result, " "))
}
