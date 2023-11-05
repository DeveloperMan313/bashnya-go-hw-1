package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < n-1; i++ {
		a, b := arr[i], arr[n-1]
		arr[n-1] = a
		arr[i] = b
	}
	str := fmt.Sprint(arr)
	fmt.Println(str[1 : len(str)-1])
}
