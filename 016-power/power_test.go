/*
@Time : 2019-11-17 14:17
@Author : lfn
@File : power_test
*/

package _16_power

import (
	"math"
	"testing"
)

func TestPower(t *testing.T) {
	tests := map[string]struct{
		base float64
		exponent int
		want float64
	}{
		"0":{1.1, 2, 1.21},
		"1":{1, 0, 1},
		"2":{2, -1, 0.5},
		"3":{0.5, 2, 0.25},
		"4":{0.5, -1, 2},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := power(tt.base, tt.exponent)
			if math.Abs(tt.want - got) > 1E-6 {
				t.Errorf("want=%v, got=%v\n", tt.want, got)
			}
		})
	}
}
