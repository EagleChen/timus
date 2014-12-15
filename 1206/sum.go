package main

import "fmt"
import "math/big"

func main() {
	var K int64
	fmt.Scan(&K)

	result := big.NewInt(1)
	result = result.Mul(big.NewInt(36), result.Exp(big.NewInt(55), big.NewInt(K-1), nil))

	fmt.Println(result)
}
