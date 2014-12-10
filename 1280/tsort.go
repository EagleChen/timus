package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func main() {
	var N, M, s, u int
	in := bufio.NewReader(os.Stdin)
	text, _ := in.ReadString('\n')
	fields := strings.Fields(text)
	N, _ = strconv.Atoi(fields[0])
	M, _ = strconv.Atoi(fields[1])

	// column 0 records the coutns
	depends := make([][]int, N+1)
	for i := 1; i <= N; i++ {
		depends[i] = make([]int, N+1)
	}

	for i := 0; i < M; i++ {
		text, _ = in.ReadString('\n')
		fields = strings.Fields(text)
		s, _ = strconv.Atoi(fields[0])
		u, _ = strconv.Atoi(fields[1])

		if depends[u][s] == 0 {
			depends[u][s] = 1
			depends[u][0]++
		}
	}

	answer := make([]int, N)
	for i := 0; i < N; i++ {
		text, _ = in.ReadString('\n')
		fields = strings.Fields(text)
		for k, v := range fields {
			answer[k], _ = strconv.Atoi(v)
		}

		if depends[answer[i]][0] == 0 {
			for j := 1; j <= N; j++ {
				if depends[j][answer[i]] == 1 {
					depends[j][0]--
				}
			}
		} else {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
