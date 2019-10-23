package main

import "fmt"

func main()  {
	matrix := [][]int{
		{1, 5, 11, 15},
		{2, 6, 12, 16},
		{3, 7, 13, 17},
		{4, 8, 14, 18},
	}
	exists, pos := ifExists(matrix, 16)
	fmt.Println("exists: ", exists)
	if exists {
		fmt.Println("pos: ", pos)
	}
}

func ifExists(arr [][]int, k int) (bool, []int) {
	rows, cols := len(arr), len(arr[0])
	i, j := 0, cols-1
	for i < rows && j >= 0 {
		if arr[i][j] == k {
			return true, []int{i, j}
		} else if arr[i][j] > k {
			j--
		} else {
			i++
		}
	}
	return false, []int{-1, -1}
}