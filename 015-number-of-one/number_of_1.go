/*
@Time : 2019-11-16 16:12
@Author : lfn
@File : number_of_1
*/

package _15_number_of_one

func numberOf1(n int) int {
	var count int
	for n != 0 {
		count++
		n = (n-1) & n
	}
	return count
}
