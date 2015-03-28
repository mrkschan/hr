package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var cases int

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Sscan(text, &cases)

	for i := 0; i < cases; i++ {
		line1, _ := reader.ReadString('\n')
		line1 = strings.TrimRight(line1, "\n")
		line2, _ := reader.ReadString('\n')
		line2 = strings.TrimRight(line2, "\n")

		chars := make([]bool, 26)
		for _, c := range line1 {
			chars[c-'a'] = true
		}

		found := false
		for _, c := range line2 {
			if chars[c-'a'] {
				found = true
				break
			}
		}

		if found {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
