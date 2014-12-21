package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	// preparation
	chars := make([]rune, len(str)*2+1)
	for i, v := range str {
		chars[2*i] = '#'
		chars[2*i+1] = v
	}
	length := len(chars)
	chars[length-1] = '#'

	// calculate
	radius := make([]int, length)
	center := 0
	last := 0
	for i := 1; i < length; i++ {
		if i <= last {
			radius[i] = radius[center*2-i]
			if radius[i]+i > last {
				radius[i] = last - i
				continue
			} else if radius[i]+i < last {
				continue
			}
		}

		// compare
		for j := radius[i] + 1; ; j++ {
			if !(i-j >= 0 && i+j < length && chars[i-j] == chars[i+j]) {
				radius[i] = j - 1
				break
			}
		}

		center = i
		last = i + radius[i]
	}

	max := 0
	// find longest
	for i := 0; i < length; i++ {
		if radius[i] > max {
			max = radius[i]
			center = i
		}
	}

	for i, end := center-max, center+max; i <= end; i++ {
		if chars[i] != '#' {
			fmt.Print(string(chars[i]))
		}
	}
}
