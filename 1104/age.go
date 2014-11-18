package main
import "fmt"
import "bufio"
import "os"
import "math"

func main() {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    text, _ := in.ReadString(0)

    dict := map[byte]int {'0':0, '1':1, '2':2, '3':3, '4':4, '5':5, '6':6, '7':7, '8':8, '9':9, 'A':10, 'B':11, 'C':12, 'D':13, 'E':14, 'F':15, 'G':16, 'H':17, 'I':18, 'J':19, 'K':20, 'L':21, 'M':22, 'N':23, 'O':24, 'P':25, 'Q':26, 'R':27, 'S':28, 'T':29, 'U':30, 'V':31, 'W':32, 'X':33, 'Y':34, 'Z':35}
    least := 1
    sum := 0
    for i := len(text) - 2; i >= 0; i-- {
        c := text[i]
        if _, ok := dict[c]; ok {
            if dict[c] > least {
                least = dict[c]
            }
            sum += dict[c]
        }
    }

    for i := least; i < 36; i++ {
        if math.Mod(float64(sum), float64(i)) == 0.0 {
            fmt.Fprintf(out, "%d\n", i+1)
            out.Flush()
            return
        }
    }

    fmt.Fprintf(out, "%s\n", "No solution.")
    out.Flush()
}
