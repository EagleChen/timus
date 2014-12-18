package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(bytes []byte, begin int, end int) {
	for ; begin < end; begin, end = begin+1, end-1 {
		bytes[begin], bytes[end] = bytes[end], bytes[begin]
	}
}

func isLatin(b byte) bool {
	if (b >= 65 && b <= 90) || (b >= 97 && b <= 122) {
		return true
	} else {
		return false
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	text, _ := in.ReadString(0)
	bytes := []byte(text)

	Word := false
	begin := 0
	length := len(bytes)
	for i := 0; i < length; i++ {
		if Word {
			if !isLatin(bytes[i]) {
				reverse(bytes, begin, i-1)
				Word = false
			}
		} else {
			if isLatin(bytes[i]) {
				begin = i
				Word = true
			}
		}
	}

	if Word {
		reverse(bytes, begin, length-1)
	}

	fmt.Fprintf(out, string(bytes))
	out.Flush()
}
