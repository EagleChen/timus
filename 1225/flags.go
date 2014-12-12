package main

import "fmt"

var remember [100]uint64

func count(slots int) uint64 {
	if remember[slots] > 0 {
		return remember[slots]
	}

	if slots == 0 {
		remember[slots] = 1
	} else if slots == 1 {
		remember[slots] = 2
	} else {
		remember[slots] = count(slots-2) + count(slots-1)
	}
	return remember[slots]
}

func main() {
	var N int
	fmt.Scan(&N)

	slots := N - 2
	if slots >= 2 {
		fmt.Println(2 * count(slots))
	} else if slots == 1 {
		fmt.Println(4)
	} else {
		fmt.Println(2)
	}
}
