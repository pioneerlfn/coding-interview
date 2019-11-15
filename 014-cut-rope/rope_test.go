/*
@Time : 2019-11-15 18:14
@Author : lfn
@File : rope_test
*/

package _14_cut_rope

import (
	"reflect"
	"testing"
)

func TestCutRope(t *testing.T) {
	tests := map[string]struct{
		length int
		product int
	}{
		"0":{2, 1},
		"1":{3, 2},
		"2":{4, 4},
		"3":{5, 6},
		"4":{6, 9},
		"5":{8, 18},

	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := cutRope(tt.length)
			if !reflect.DeepEqual(tt.product, got) {
				t.Errorf("want=%v, got=%v\n", tt.product, got)
			}
		})
	}
}


