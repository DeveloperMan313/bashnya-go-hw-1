package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d\n%d\n%d", &a, &b, &c)
	if a < b+c && b < a+c && c < a+b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
