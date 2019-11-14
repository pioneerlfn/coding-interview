/*
@Time : 2019-11-13 22:01
@Author : lfn
@File : min_test
*/

package _11_min_rotate_arr

import (
	"reflect"
	"testing"
)

func TestMinInRotate(t *testing.T) {
	tests := map[string]struct{
		arr []int
		want int
	}{
		"1":{[]int{3,4,5,1,2},1},
		"2":{[]int{4,3}, 3},
		"3":{[]int{1,2,3}, 1},
		"4":{[]int{0,1,1,1,1}, 0},
		"5":{[]int{1,1,1,1,0,1}, 0},
		"6":{[]int{1,1,1,0,1,1}, 0},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := minInRotate(tt.arr)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want=%v, got=%v\n", tt.want, got)
			}
		})
	}
}
