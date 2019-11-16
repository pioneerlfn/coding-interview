/*
@Time : 2019-11-16 23:15
@Author : lfn
@File : power
*/

package _16_power

import (
	"errors"
)

func power(base float64, exponet int) (float64, error) {
	if base == 0 && exponet < 0 {
		return 0, errors.New("invalid input")
	}
	if exponet < 0 {
		exponet = -exponet
	}
	res := powerWithAbsExponet(base, uint(exponet))
	if exponet < 0 {
		res = 1 / res
	}
	return res, nil
}

func powerWithAbsExponet(base float64, exponent uint) float64 {
	res := 1.0
	for i := uint(1); i <= exponent; i++ {
		res *= base
	}
	return res
}



