/*
@Time : 2019-11-16 16:14
@Author : lfn
@File : count_test
*/

package _15_number_of_one

import (
	"reflect"
	"testing"
)

func TestNumberOf1(t *testing.T) {
	tests := map[string]struct{
		n int
		want int
	}{
		"0":{0, 0},
		"1":{1, 1},
		"2":{2, 1},
		"3":{9, 2},
		"4":{-1, 64},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := numberOf1(tt.n)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want=%v, got=%v\n", tt.want, got)
			}
		})
	}
}
