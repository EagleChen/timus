package main

import (
	"bytes"
	"fmt"
)

func a(pre string, num int) stirng {

}

func s(num int) string {
	var aStr = ""
	var result bytes.Buffer
	for i := 1; i < num; i++ {
		result.WriteString("(")
	}
	for i := 1; i <= num; i++ {
		aStr = a(aStr, i)
	}
}

func main() {
	var num int
	fmt.Scan(&num)

	fmt.Println(s(num))
}
