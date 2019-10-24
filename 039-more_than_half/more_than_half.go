package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 2, 6, 2, 9, 2, 2}
	it := findMoreThanHalf(arr)
	fmt.Println(it)

}

func findMoreThanHalf(arr []int) int {
	low, high := 0, len(arr)-1
	k := high >> 1
	num := qSelect(arr, low, high, k)
	if checkMoreThanHalf(arr, num) {
		return num
	}
	return -1
}

func qSelect(arr []int, low, high int, k int) int {
	idx := partition(arr, low, high)
	if idx+1 == k {
		return arr[k]
	} else if idx+1 < k {
		return qSelect(arr, idx+1, high, k)
	} else {
		return qSelect(arr, low, idx-1, k)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	for low < high {
		for low < high && arr[high] >= pivot {
			high--
		}
		arr[low] = arr[high]
		for low < high && arr[low] <= pivot {
			low++
		}
		arr[high] = arr[low]
		arr[low] = pivot
	}
	return low
}

func checkMoreThanHalf(arr []int, num int) bool {
	var cnt int
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			cnt++
		}
	}
	if cnt > (len(arr) >> 1) {
		return true
	}
	return false
}
