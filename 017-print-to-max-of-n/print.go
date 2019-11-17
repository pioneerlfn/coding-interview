/*
@Time : 2019-11-17 17:24
@Author : lfn
@File : print
*/

// package _17_print_to_max_of_n
package main


import (
	"fmt"
)

func print1ToMaxNumberOfNDigits(n int) {
	if n <= 0 {
		return
	}
	numbers := make([]int, n)
	for i := 0; i < 10; i++ {
		numbers[0] = i + '0'
		printRecursive(numbers, 0)
	}
}

func printRecursive(nums []int, idx int) {
	if idx == len(nums)-1 {
		print(nums)
		return
	}
	for i := 0; i < 10; i++ {
		nums[idx+1] = i + '0'
		printRecursive(nums, idx+1)
	}
}

func print(nums []int) {
	beginning0 := true
	for i := 0; i < len(nums); i++ {
		if beginning0 && nums[i] != '0' {
			beginning0 = false
		}
		if !beginning0 {
			fmt.Printf("%c", nums[i])
		}
	}
	fmt.Println()
}

func main() {
	print1ToMaxNumberOfNDigits(2)
}

