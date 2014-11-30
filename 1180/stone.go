package main

import "fmt"

func main() {
	var N string
	fmt.Scan(&N)

	num := 0
	for i := len(N) - 1; i >= 0; i-- {
		num += int(N[i] - '0')
	}

	remainder := num % 3
	if remainder == 0 {
		fmt.Println(2)
	} else {
		fmt.Println(1)
		fmt.Println(remainder)
	}
}
