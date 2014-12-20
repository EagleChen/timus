package main

import (
	"fmt"
	"math/big"
)

func main() {
	var num string
	fmt.Scan(&num)

	i := new(big.Int)
	i.SetString(num, 10)
	div := big.NewInt(int64(7))
	fmt.Println(i.Mod(i, div))
}
