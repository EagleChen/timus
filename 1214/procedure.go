package main

import "fmt"

func main() {
	var x, y int

	fmt.Scan(&x)
	fmt.Scan(&y)

	if x > 0 && y > 0 {
		if (x+y)%2 != 0 {
			x, y = y, x
		}
	}

	fmt.Printf("%d %d\n", x, y)
}
