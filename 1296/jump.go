package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(nums []int) int {
	max := 0
	tail := 0

	for i := len(nums) - 1; i >= 0; i-- {
		sum := nums[i] + tail
		if sum > max {
			max = sum
		}
		if sum > 0 {
			tail = sum
		} else {
			tail = 0
		}
	}
	return max
}

func main() {
	in := bufio.NewReader(os.Stdin)
	text, _ := in.ReadString(0)

	fields := strings.Fields(text)
	size, _ := strconv.Atoi(fields[0])

	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i], _ = strconv.Atoi(fields[i+1])
	}

	fmt.Println(max(nums))
}
