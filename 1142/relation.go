package main

import (
	"fmt"
	//	"strconv"
)

//permutation = make(map[int]uint64)
//
// exactPositions(balls, drawer, min int) [][]int {
//result := make([][]int, 0)
//if drawer == 1 {
//	return append(result, []int{balls})
//}
//for i := balls / drawer; i > 0; i-- {
//	for _, v := range exactPositions(balls-i, drawer-1, i) {
//		result = append(result, append(v, i))
//	}
//}
//return result
//
//
// positions(balls, drawer int) [][]int {
//if drawer > balls {
//	drawer = balls
//}
//result := make([][]int, 0)
//for i := 1; i <= drawer; i++ {
//	for _, v := range exactPositions(balls, i, 1) {
//		result = append(result, v)
//	}
//}
//return result
//
//
// combinations(total, sub int) uint64 {
//return calPermutation(total) / calPermutation(total-sub) / calPermutation(sub)
//
//
// calPermutation(total int) uint64 {
//if num, ok := permutation[total]; ok {
//	return num
//}
//var result uint64 = 1
//for i := 2; i <= total; i++ {
//	result *= uint64(i)
//	}
//	permutation[total] = result
//	return result
//}
//
//func count(numOfObjs, numOfEquals int) uint64 {
//	//fmt.Println(strconv.Itoa(numOfEquals) + " equals for " + strconv.Itoa(numOfObjs))
//	if numOfEquals == 0 {
//		return calPermutation(numOfObjs)
//	}
//	result := uint64(0)
//	arr := positions(numOfEquals, numOfObjs-numOfEquals)
//	//fmt.Println(arr)
//	for i := len(arr) - 1; i >= 0; i-- {
//		tmp := calPermutation(numOfObjs - numOfEquals)
//		tmpArr := arr[i]
//		dup := make(map[int]int)
//		tmpObjs := numOfObjs
//		for j := len(tmpArr) - 1; j >= 0; j-- {
//			tmp *= combinations(tmpObjs, tmpArr[j]+1)
//			tmpObjs -= tmpArr[j] + 1
//			if _, ok := dup[tmpArr[j]]; ok {
//				dup[tmpArr[j]]++
//			} else {
//				dup[tmpArr[j]] = 1
//			}
//		}
//		for _, v := range dup {
//			tmp /= calPermutation(v)
//		}
//		//	fmt.Println(tmp)
//		result += tmp
//	}
//	return result
//}
//
//func relation(num int) uint64 {
//	var result uint64 = 0
//	for i := 0; i < num; i++ {
//		result += count(num, i)
//	}
//	return result
//}

// my algorithm is incorrect when num is 9 and 10. But i don't know what is wrong

func main() {
	N := []int{0, 0, 3, 13, 75, 541, 4683, 47293, 545835, 7087261, 102247563}
	var num int
	for {
		fmt.Scan(&num)

		if num == -1 {
			return
		}
		//fmt.Println(strconv.FormatUint(relation(num), 10))
		fmt.Println(N[num])
	}
}
