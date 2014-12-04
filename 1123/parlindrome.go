package main

import "fmt"

func increase(arr []byte, middle int) {
	for i := middle; i >= 0; i-- {
		if arr[i] == '9' {
			arr[i] = '0'
		} else {
			arr[i] += 1
			break
		}
	}
}

func main() {
	var str string

	fmt.Scan(&str)

	arr := []byte(str)
	middle := -1
	post := -1
	pre := -1
	if len(arr)%2 != 0 {
		middle = len(arr) / 2
		post = middle + 1
		pre = middle - 1
	} else {
		middle = len(arr)/2 - 1
		post = middle + 1
		pre = middle
	}

	for i, j := pre, post; i >= 0; i, j = i-1, j+1 {
		if arr[j] > arr[i] {
			increase(arr, middle)
			break
		} else if arr[j] < arr[i] {
			break
		}
	}

	for ; pre >= 0; pre, post = pre-1, post+1 {
		arr[post] = arr[pre]
	}

	fmt.Println(string(arr))
}
