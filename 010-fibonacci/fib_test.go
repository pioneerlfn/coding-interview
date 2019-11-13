/*
@Time : 2019-11-13 21:38
@Author : lfn
@File : fib_test
*/

package _10_fibonacci

import (
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	tests := map[string]struct{
		n int
		want int
	}{
		"1":{0,1},
		"2":{1, 1},
		"3":{2, 2},
		"4":{3, 3},
		"5":{4, 5},
		"6":{10, 89},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := fib(tt.n)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want=%v, got=%v\n", tt.want, got)
			}
		})
	}
}

