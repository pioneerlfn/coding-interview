/*
@Time : 2019-11-13 21:55
@Author : lfn
@File : min
*/

package _11_min_rotate_arr

func minInRotate(arr []int) int {
	low, high := 0, len(arr)-1
	if arr[high] > arr[low] {
		return arr[low]
	}
	mid := low + (high-low) >> 1
	if arr[low] == arr[high] && arr[low] == arr[mid] {
		return findMinLinear(arr)
	}
	return min(arr, 0, len(arr)-1)
}

func min(arr []int, low, high int) int {
	for low < high {
		if high - low == 1 {
			break
		}
		mid := low + (high-low)>>1
		if arr[low] <= arr[mid] {
			low = mid
		} else {
			high = mid
		}
	}
	return arr[high]
}

func findMinLinear(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}