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
	list []int
}

func combinable(x, y, lSquare int) bool {
	return x*x+y*y <= lSquare
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

	v1 := vec{0, 0, make([]int, 0)}
	v2 := vec{0, 0, make([]int, 0)}
	for i := 0; i < num; i++ {
		xY, _ := in.ReadString('\n')
		strs := strings.Split(strings.TrimSpace(xY), " ")
		x, _ := strconv.Atoi(strs[0])
		y, _ := strconv.Atoi(strs[1])

		x1 := v1.x + x
		y1 := v1.y + y
		if combinable(x1, y1, lSquare) {
			v1.x = x1
			v1.y = y1
			v1.list = append(v1.list, i)
		} else {
			x2 := v2.x + x
			y2 := v2.y + y
			if combinable(x2, y2, lSquare) {
				v2.x = x2
				v2.y = y2
				v2.list = append(v2.list, i)
			} else {
				v1.x += v2.x
				v1.y += v2.y
				v1.list = append(v1.list, v2.list...)

				v2.x, v2.y = 0, 0
				v2.list = v2.list[:0]
			}
		}
	}

	fmt.Fprintln(out, "YES")
	xFinal := v1.x + v2.x
	yFinal := v1.y + v2.y
	if combinable(xFinal, yFinal, lSquare) {
		for i := 0; i < num; i++ {
			fmt.Fprint(out, "+")
		}
		fmt.Fprintf(out, "\n")
	} else {
		result := make([]string, num)
		for i, _ := range result {
			result[i] = "+"
		}
		for _, v := range v1.list {
			result[v] = "-"
		}
		fmt.Fprintln(out, strings.Join(result, ""))
	}
	out.Flush()
}
