package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int64
	fmt.Scanf("%d\n%d", &a, &b)
	fmt.Println(math.Sqrt(float64(a*a + b*b)))
}
