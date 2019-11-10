/*
@Time : 2019-11-10 19:27
@Author : lfn
@File : singleton
*/

package _2_singleton

import (
	"sync"
)

type Singleton map[string]string

var (
	once sync.Once
	instance Singleton
)

func New() Singleton {
	once.Do(func() {
		instance = make(Singleton)
	})
	return instance
}



