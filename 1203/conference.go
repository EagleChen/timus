package main

import "fmt"
import "sort"
import "bufio"
import "os"
import "strings"
import "strconv"

type span struct {
	start int
	end   int
}

type spans []span

func (s spans) Len() int           { return len(s) }
func (s spans) Less(i, j int) bool { return s[i].end < s[j].end }
func (s spans) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	var N, start, end int
	fmt.Scan(&N)

	in := bufio.NewReader(os.Stdin)
	text, _ := in.ReadString(0)
	fields := strings.Fields(text)
	conferences := make(spans, N)
	for i := 0; i < N; i++ {
		start, _ = strconv.Atoi(fields[2*i])
		end, _ = strconv.Atoi(fields[2*i+1])
		conferences[i] = span{start, end + 1}
	}

	sort.Sort(conferences)

	count := 1
	max := conferences[0].end
	for i := 1; i < N; i++ {
		if conferences[i].start >= max {
			count++
			max = conferences[i].end
		}
	}

	fmt.Println(count)
}
