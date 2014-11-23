package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func a(pre string, num int) string {
	if num == 1 {
		return "sin(1)"
	} else {
		end := len(pre) - num + 1
		result := bytes.NewBufferString(pre[:end])
		if num%2 == 0 {
			result.WriteString("-")
		} else {
			result.WriteString("+")
		}

		result.WriteString("sin(")
		result.WriteString(strconv.Itoa(num))
		for i := 0; i < num; i++ {
			result.WriteString(")")
		}
		return result.String()
	}
}

func s(num int) string {
	var aStr = ""
	var result bytes.Buffer
	for i := 1; i < num; i++ {
		result.WriteString("(")
	}
	for i := 1; i <= num; i++ {
		aStr = a(aStr, i)
		result.WriteString(aStr + "+" + strconv.Itoa(num+1-i) + ")")
	}
	result.Truncate(result.Len() - 1)
	return result.String()
}

func main() {
	var num int
	fmt.Scan(&num)

	fmt.Println(s(num))
}
