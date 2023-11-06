package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	set := make(map[string]struct{})
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nums := strings.Fields(scanner.Text())
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = struct{}{}
	}
	fmt.Println(len(set))
}
