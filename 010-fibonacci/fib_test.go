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
		"1":{0,0},
		"2":{1, 1},
		"3":{2, 1},
		"4":{3, 2},
		"5":{4, 3},
		"6":{10, 55},
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

