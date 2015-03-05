package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func find_left(nums []int) (int, int, bool) {
	min := 0

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			if nums[i] >= min {
				return i - 1, min, true
			} else {
				return -1, -1, false
			}
		}
		min = nums[i-1]
	}

	return -1, -1, false
}

func find_right(nums []int, left int, min int) (int, bool) {
	max := 1000000
	for i := len(nums) - 1; i > left; i-- {
		if nums[i-1] > nums[i] {
			for j := i; j < len(nums)-1 && nums[j] == nums[j+1]; j++ {
				i = j + 1
			}
			if nums[i] >= min && nums[left] <= max {
				return i, true
			} else {
				return -1, false
			}
		}
		if nums[i-1] < nums[i] {
			max = nums[i]
		}
	}

	return -1, false
}

func is_asc(nums []int, l int, r int) bool {
	for i := l; i < r; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}

	return true
}

func is_desc(nums []int, l int, r int) bool {
	for i := l; i < r; i++ {
		if nums[i] < nums[i+1] {
			return false
		}
	}

	return true
}

func main() {
	var size int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	size, _ = strconv.Atoi(scanner.Text())

	nums := make([]int, size)
	var num int
	for i := 0; i < size; i++ {
		scanner.Scan()
		num, _ = strconv.Atoi(scanner.Text())

		nums[i] = num
	}

	// 1 5 3 4 4 2 6 | swap(2,6)
	// = + - + = - +
	//   ^       ^
	// 1 5 4 4 3 2 6 | reverse(2,6)
	// = + - = - - +
	//   ^       ^
	// 3 6 4 3 5 5 2 6 | no
	// = + - - + = - +
	//   ^         x

	if len(nums) < 2 {
		fmt.Println("yes")
		return
	}

	l, min, found := find_left(nums)
	if !found {
		if is_asc(nums, 0, len(nums)-1) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
		return
	}

	r, found := find_right(nums, l, min)
	if !found {
		fmt.Println("no")
		return
	}

	if is_asc(nums, l+1, r-1) {
		fmt.Println("yes")
		fmt.Printf("swap %d %d\n", l+1, r+1)
	} else if is_desc(nums, l+1, r-1) {
		fmt.Println("yes")
		fmt.Printf("reverse %d %d\n", l+1, r+1)
	} else {
		fmt.Println("no")
	}
}
