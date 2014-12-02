package main

import "fmt"

func main() {
	var N int
	str := ""
	tmp := ""
	fmt.Scan(&N)
	if N == 1 {
		fmt.Println(0)
		return
	}

	for len(str) < N {
		fmt.Scan(&tmp)
		str += tmp
	}

	// naive method
	// bytes := []byte(str)
	//	visited := make(map[string]bool)
	//
	//	total := 0
	//	for {
	//		if visited[str] {
	//			fmt.Println("NO")
	//			return
	//		}
	//		count := 0
	//		visited[str] = true
	//		for i := 1; i < N; {
	//			if bytes[i-1] == '>' && bytes[i] == '<' {
	//				count += 1
	//				bytes[i-1] = '<'
	//				bytes[i] = '>'
	//				i += 2
	//			} else {
	//				i++
	//			}
	//		}
	//		if count == 0 {
	//			fmt.Println(total)
	//			return
	//		} else {
	//			total += count
	//			str = string(bytes[:N])
	//		}
	//	}

	index := N - 1
	count := 0
	for i := N - 1; i >= 0; i-- {
		if str[i] == '>' {
			count += index - i
			index--
		}
	}
	fmt.Println(count)
}
