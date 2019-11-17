/*
@Time : 2019-11-16 23:15
@Author : lfn
@File : power
*/

package _16_power

import (
	"errors"
)

func power(base float64, exponent int) (float64, error) {
	if base == 0 && exponent < 0 {
		return 0, errors.New("invalid input")
	}
	var negative bool
	if exponent < 0 {
		negative = true
		exponent = -exponent
	}
	res := powerWithAbsExponent(base, uint(exponent))
	if negative {
		res = 1 / res
	}
	return res, nil
}

func powerWithAbsExponent(base float64, exponent uint) float64 {
	if exponent == 0 {
		return 1.0
	}
	if exponent == 1 {
		return base
	}
	/*for i := uint(1); i <= exponent; i++ {
		res *= base
	}*/
	if exponent % 2 == 1 {
		return base * powerWithAbsExponent(base*base, exponent / 2)
	}
	return powerWithAbsExponent(base*base, exponent / 2)

}



