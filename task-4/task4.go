package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n*n)
	for i := 0; i < n*n; i++ {
		fmt.Scan(&arr[i])
	}
	symmetric := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[i*n+j] != arr[j*n+i] {
				symmetric = false
				goto endloop
			}
		}
	}
endloop:
	if symmetric {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
