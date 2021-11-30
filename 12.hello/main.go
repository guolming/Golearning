package main

import (
	"fmt"
	"math"
)

func main() {
	ver, rct := 34, 89.73442
	var sc = ver * int(rct)
	fmt.Println(ver + int(rct))
	fmt.Println(sc)

	var int6 uint = math.MaxUint64
	fmt.Println(int6)
	var int7 int = int(int6)
	fmt.Println(int7)
}
