package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vec struct {
	x    int
	y    int
	list []string
}

func combine(list1, list2 []string, direction string) {
	num := len(list1)
	if direction == "+" {
		for j := 0; j < num; j++ {
			if list2[j] != "" {
				list1[j] = list2[j]
				list2[j] = ""
			}
		}
	} else {
		for j := 0; j < num; j++ {
			switch list2[j] {
			case "+":
				list1[j] = "-"
				list2[j] = ""
			case "-":
				list1[j] = "+"
				list2[j] = ""
			}
		}
	}
}

func combinable(x1, y1, x2, y2, lSquare int) (int, int, string, bool) {
	x := x1 + x2
	y := y1 + y2
	if x*x+y*y <= lSquare {
		return x, y, "+", true
	} else {
		x := x1 - x2
		y := y1 - y2
		if x*x+y*y <= lSquare {
			return x, y, "-", true
		} else {
			return -1, -1, "", false
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	inputNum, _ := in.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(inputNum))
	inputL, _ := in.ReadString('\n')
	l, _ := strconv.Atoi(strings.TrimSpace(inputL))
	lSquare := l * l

	if num == 1 {
		fmt.Fprintln(out, "YES")
		fmt.Fprintln(out, "+")
		out.Flush()
		return
	}

	v1 := vec{0, 0, make([]string, num)}
	v2 := vec{0, 0, make([]string, num)}
	for i := 0; i < num; i++ {
		xY, _ := in.ReadString('\n')
		strs := strings.Split(strings.TrimSpace(xY), " ")
		x, _ := strconv.Atoi(strs[0])
		y, _ := strconv.Atoi(strs[1])

		if resX, resY, direction, ok := combinable(v1.x, v1.y, x, y, lSquare); ok {
			v1.x = resX
			v1.y = resY
			v1.list[i] = direction
		} else {
			if resX, resY, direction, ok := combinable(v2.x, v2.y, x, y, lSquare); ok {
				v2.x = resX
				v2.y = resY
				v2.list[i] = direction
			} else {
				resX, resY, direction, _ := combinable(v1.x, v1.y, v2.x, v2.y, lSquare)
				v1.x = resX
				v1.y = resY
				combine(v1.list, v2.list, direction)

				v2.x, v2.y = x, y
				v2.list[i] = "+"
			}
		}
	}

	fmt.Fprintln(out, "YES")
	_, _, direction, _ := combinable(v1.x, v1.y, v2.x, v2.y, 2*lSquare)
	combine(v1.list, v2.list, direction)
	fmt.Fprintln(out, strings.Join(v1.list, ""))
	out.Flush()
}
