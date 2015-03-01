package main

import (
	"fmt"
)

func rightCircularRotate(head *int, size int) {
	if *head == 0 {
		*head = size - 1
	} else {
		*head = *head - 1
	}
}

func materializedIdx(idx int, head int, size int) int {
	return (head + idx) % size
}

func main() {
	var size, k, q int
	fmt.Scanln(&size, &k, &q)

	array := make([]int, size)
	for i := 0; i < size; i++ {
		var x int
		fmt.Scan(&x)

		array[i] = x
	}

	head := 0
	for i := 0; i < k; i++ {
		rightCircularRotate(&head, size)
	}

	for i := 0; i < q; i++ {
		var x, y int
		fmt.Scan(&x)

		y = materializedIdx(x, head, size)
		fmt.Println(array[y])
	}
}
