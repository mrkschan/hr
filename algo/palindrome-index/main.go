package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

func is_palindrome(text string) bool {
	i := 0
	j := len(text) - 1

	for i < j {
		if text[i] != text[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func find_palindrome(text string) int {
	idx := -1
	i := 0
	j := len(text) - 1

	for i < j && text[i] == text[j] {
		i++
		j--
	}

	if i >= j {
		return idx
	}

	if is_palindrome(text[i:j]) {  // Skip J
		return j
	} else {
		return i
	}
}

func main() {
	var text string
	reader := bufio.NewReader(os.Stdin)

	var cases int

	text, _ = reader.ReadString('\n')
	fmt.Sscan(text, &cases)

	for i := 0; i < cases; i++ {
		text, _ = reader.ReadString('\n')

		idx := find_palindrome(strings.TrimRight(text, "\n"))
		fmt.Println(idx)
	}
}
