package main

import "fmt"

type Num struct {
	index int
	value int
}

func addNum(values []Num, num Num, window int) []Num {
	if len(values) == 0 || window == 1 {
		return []Num{num}
	}

	start := 0
	if num.index-values[0].index == window {
		start = 1
	}

	end := len(values) - 1
	for ; end >= start && num.value >= values[end].value; end-- {
	}

	if end < start {
		return []Num{num}
	}

	return append(values[start:end+1], num)
}

func main() {
	var window int

	fmt.Scan(&window)

	values := []Num{}
	input := 0
	i := 0
	for ; i < window; i++ {
		fmt.Scan(&input)
		values = addNum(values, Num{i, input}, window)
	}

	for ; ; i++ {
		fmt.Println(values[0].value)

		fmt.Scan(&input)
		if input == -1 {
			break
		}
		values = addNum(values, Num{i, input}, window)
	}
}
