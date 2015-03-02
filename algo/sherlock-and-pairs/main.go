package main

import (
	"bufio"
	"fmt"
	"os"
)

// func nPr(n int, r int) int {
// 	x := 1
// 	for i := 0; i <= n; i++ {
// 		if i > (n - r) {
// 			x = x * i
// 		}
// 	}
// 	return x
// }

func nP2(n int64) int64 {
	return n * (n - 1)
}

func main() {
	var cases, size int
	var total int64

	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &cases)

	for i := 0; i < cases; i++ {
		fmt.Fscan(io, &size)

		freq := make(map[int]int64)
		var x int
		for j := 0; j < size; j++ {
			fmt.Fscan(io, &x)
			freq[x] += 1
		}

		total = 0
		for _, c := range freq {
			if c > 1 {
				total += nP2(c)
			}
		}

		fmt.Println(total)
	}
}
