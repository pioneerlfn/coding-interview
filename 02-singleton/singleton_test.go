/*
@Time : 2019-11-10 19:30
@Author : lfn
@File : singleton_test
*/

package _2_singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	s := New()
	s["this"] = "that"
	s2 := New()
	want := "that"
	if got := s2["this"]; got != want {
		t.Errorf("want = %s, got = %s\n", want, got)
	}
}
