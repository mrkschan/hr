package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	counter := make([]int, 10001)

	var size int
	var n int

	scanner.Scan()
	size, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < size; i++ {
		scanner.Scan()
		n, _ = strconv.Atoi(scanner.Text())

		counter[n]++
	}

	min := 10001
	scanner.Scan()
	size, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < size; i++ {
		scanner.Scan()
		n, _ = strconv.Atoi(scanner.Text())

		if n < min {
			min = n
		}

		counter[n]--
	}

	missing := []string{}
	for i := 0; i < 101; i++ {
		if counter[min + i] != 0 {
			missing = append(missing, strconv.Itoa(min + i))
		}
	}

	fmt.Println(strings.Join(missing, " "))
}
