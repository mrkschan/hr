package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var size, diff, n int

	scanner.Scan()
	size, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	diff, _ = strconv.Atoi(scanner.Text())

	nums := make(map[int]bool)
	for i := 0; i < size; i++ {
		scanner.Scan()
		n, _ = strconv.Atoi(scanner.Text())

		nums[n] = true
	}

	pairs := 0
	for k, _ := range nums {
		_, ok := nums[k+diff]
		if ok {
			pairs++
		}
	}

	fmt.Println(pairs)
}
