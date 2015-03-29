package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func next_permutation(text string) (bool, string) {
	// Ref - http://stackoverflow.com/a/11483392/433662
	// Swap if all stuffs on the right is in DESC
	bytes := []byte(text)

	for j := len(bytes) - 1; j > 0; j-- {
		if bytes[j-1] < bytes[j] {
			p := j - 1

			// Find k, the next smallest element that's larger than p
			k := len(bytes) - 1
			for ; bytes[p] >= bytes[k]; k-- {
			}

			bytes[k], bytes[p] = bytes[p], bytes[k]
			lpart, rpart := bytes[:j], bytes[j:]

			// Reverse rpart, rpart changed from DESC to ASC
			for x, y := 0, len(rpart)-1; x < y; x, y = x+1, y-1 {
				rpart[x], rpart[y] = rpart[y], rpart[x]
			}

			next := []byte{}
			next = append(next, lpart...)
			next = append(next, rpart...)

			return true, string(next)
		}
	}

	return false, text
}

func main() {
	var cases int

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Sscan(text, &cases)

	for i := 0; i < cases; i++ {
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\n")

		mutated, next := next_permutation(text)

		if mutated {
			fmt.Println(next)
		} else {
			fmt.Println("no answer")
		}
	}
}
