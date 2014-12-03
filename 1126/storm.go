package main

import (
	"fmt"
	"sort"
)

func main() {
	var count int

	fmt.Scan(&count)

	values := []int{}
	num := 0
	for i := 0; i < count; i++ {
		fmt.Scan(&num)
		values = append(values, num)
	}

	compare := make([]int, count)
	for {
		copy(compare, values)
		sort.Ints(compare)
		fmt.Println(compare[count-1])

		fmt.Scan(&num)
		if num == -1 {
			break
		}
		values = append(values[1:count], num)
	}
}
