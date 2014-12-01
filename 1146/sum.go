package main

import "fmt"

// at least one element
func maxSum(arr []int) int {
	max := -128
	tail := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	for _, v := range arr {
		sum := tail + v
		if sum > max {
			max = sum
		}

		if sum > 0 {
			tail = sum
		} else {
			tail = 0
		}
	}

	return max
}

func main() {
	var N int
	fmt.Scan(&N)

	arr := make([][]int, N)
	for i := 0; i < N; i++ {
		arr[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

	sum := make([]int, N)
	max := -2147483648
	for i := 0; i < N; i++ {
		copy(sum, arr[i])
		result := maxSum(sum)
		if result > max {
			max = result
		}
		for j := i + 1; j < N; j++ {
			for k := 0; k < N; k++ {
				sum[k] += arr[j][k]
			}
			result := maxSum(sum)
			if result > max {
				max = result
			}
		}
	}

	fmt.Println(max)
}
