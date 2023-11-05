package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	set := make(map[string]bool)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nums := strings.Fields(scanner.Text())
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = true
	}
	fmt.Println(len(set))
}
