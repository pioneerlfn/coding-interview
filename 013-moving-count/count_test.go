/*
@Time : 2019-11-15 12:25
@Author : lfn
@File : count_test
*/

package _13_moving_count

import (
	"reflect"
	"testing"
)

func TestMovingCount(t *testing.T) {
	tests := map[string]struct{
		threhold int
		rows, cols int
		want int
	}{
		"0":{5, 4, 3, 11},
		"1":{18,3, 5, 15},
		"2":{5,3, 5, 12},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := movingCount(tt.threhold, tt.rows, tt.cols)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want=%v, got=%v\n", tt.want, got)
			}
		})
	}
}

