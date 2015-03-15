package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func max(nums ...int) int {
	max := 0
	for _, i := range nums {
		if i > max {
			max = i
		}
	}
	return max
}

func main() {
	var textA, textB string
	reader := bufio.NewReader(os.Stdin)

	textA, _ = reader.ReadString('\n')
	textB, _ = reader.ReadString('\n')

	textA = strings.TrimRight(textA, "\n")
	textB = strings.TrimRight(textB, "\n")

	lcs := make([][]int, len(textA)+1)
	for i := range lcs {
		lcs[i] = make([]int, len(textB)+1)
	}

	lcs[0][0] = 0
	lcs[0][1] = 0
	lcs[1][0] = 0

	// BECD
	// BCDE
	longest := 0
	for i := 1; i <= len(textA); i++ {
		for j := 1; j <= len(textB); j++ {
			if textA[i-1] == textB[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = max(lcs[i][j-1], lcs[i-1][j])
			}

			if lcs[i][j] > longest {
				longest = lcs[i][j]
			}
		}
	}

	// DEBUG :)
	// for i := range lcs {
	// 	fmt.Println(lcs[i])
	// }

	fmt.Println(longest)
}
