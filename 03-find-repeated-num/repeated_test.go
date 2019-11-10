/*
@Time : 2019-11-10 18:55
@Author : lfn
@File : repeated_test
*/

package _3_find_repeated_num

import (
	"testing"
)

func TestRepeated(t *testing.T) {
	tests := map[string]struct{
		arr []int
		want int
	}{
		"1": {[]int{1, 2, 3, 2, 4, 5, 1},2},
		"2": {[]int{2, 4, 1, 3, 4}, 4},
		"3": {[]int{}, -1},
		"4": {[]int{1,1,1,1}, 1},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := findRepeated(tt.arr)
			// t.Log("arr: ", tt, "got: ", got)
			if got != tt.want {
				t.Errorf("want=%d, got=%d\n", tt.want, got)
			}
		})
	}
}
