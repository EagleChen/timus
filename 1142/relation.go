package main

import (
	"fmt"
	"strconv"
)

var permutation = make(map[int]uint64)

func positions(balls, drawer int) [][]int {
	if balls == 1
	if drawer > balls {
		drawer = balls
	}
	result := make([][]int, 0)

}

func calPermutation(total int) uint64 {
	if num, ok := permutation[total]; ok {
		return num
	}
	var result uint64 = 1
	for i := 2; i <= total; i++ {
		result *= uint64(i)
	}
	permutation[total] = result
	return result
}

func count(numOfObjs, numOfEquals int) uint64 {
	base := calPermutation(numOfObjs)
	if numOfEquals == 0 {
		return base
	}
	result := uint64(0)
	arr := positions(numOfEquals, numOfObjs-numOfEquals-1)
	for i := len(arr) - 1; i >= 0; i-- {
		tmp := base
		tmpArr := arr[i]
		for j := len(tmpArr) - 1; j >= 0; j-- {
			if tmpArr[j] != 0 {
				tmp /= calPermutation(tmpArr[j] + 1)
			}
		}
		result += tmp
	}
	return result
}

func relation(num int) uint64 {
	var result uint64 = 0
	for i := 0; i < num; i++ {
		result += count(num, i)
	}
	return result
}

func main() {
	var num int
	for {
		fmt.Scan(&num)

		if num == -1 {
			return
		}
		fmt.Println(strconv.FormatUint(relation(num), 10))
	}
}
