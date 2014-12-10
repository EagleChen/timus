package main

import "fmt"

func main() {
	var c, v, choice int
	fmt.Scan(&c)
	fmt.Scan(&v)

	counts := make([]int, c)
	for i := 0; i < v; i++ {
		fmt.Scan(&choice)
		counts[choice-1]++
	}

	for i := 0; i < c; i++ {
		fmt.Printf("%.2f%%\n", 100.0*float32(counts[i])/float32(v))
	}
}
