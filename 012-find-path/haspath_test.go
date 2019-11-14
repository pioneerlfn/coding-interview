/*
@Time : 2019-11-14 22:43
@Author : lfn
@File : haspath_test
*/

package _12_find_path

import (
	"reflect"
	"testing"
)

func TestHasPath(t *testing.T) {
	matrix := [][]byte{
		{'a', 'b', 't', 'g'},
		{'c', 'f', 'c', 's'},
		{'j', 'd', 'e', 'h'},
	}

	tests := map[string]struct{
		target string
		hasPath bool
	}{
		"1":{"bfce",true},
		"2":{"abfb", false},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := hasPath(matrix, tt.target)
			if !reflect.DeepEqual(tt.hasPath, got) {
				t.Errorf("want=%v, got=%v\n", tt.hasPath, got)
			}
		})
	}
}
