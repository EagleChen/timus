package main

import "fmt"

func jumpRow(row, col int) int {
	var m1, m2 int
	if row > 2 && row < 7 {
		m1 = 2
	} else {
		m1 = 1
	}

	if col > 1 && col < 8 {
		m2 = 2
	} else {
		m2 = 1
	}

	return m1 * m2
}

func handle(str string) int {
	row := int(str[0] - 'a' + 1)
	col := int(str[1] - '0')

	return jumpRow(row, col) + jumpRow(col, row)
}

func main() {
	var num int
	fmt.Scan(&num)

	var str string
	for i := 0; i < num; i++ {
		fmt.Scan(&str)
		fmt.Println(handle(str))
	}
}
