package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	mat := make([]int, n*n)
	for i := 0; i < n*n; i++ {
		fmt.Scan(&mat[i])
	}
	if matrixSymmetric(mat, n) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func matrixSymmetric(mat []int, n int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i*n+j] != mat[j*n+i] {
				return false
			}
		}
	}
	return true
}
