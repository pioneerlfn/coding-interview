/*
@Time : 2019-11-10 18:48
@Author : lfn
@File : repeated
*/

package _3_find_repeated_num


func findRepeated(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	for i := 0; i < len(arr); i++ {
		for arr[i] != i && arr[i] != arr[arr[i]]{
			arr[arr[i]], arr[i] = arr[i], arr[arr[i]]
		}
		if arr[i] == arr[arr[i]] {
			return arr[i]
		}
	}
	return -1
}
