/*
@Time : 2019-11-10 20:38
@Author : lfn
@File : space_test
*/

package _5_space_replace

import (
	"testing"
)

func TestReplaceSpace(t *testing.T) {
	tests := map[string]struct{
		source string
		want string
	}{
		"happy": {"we are happy.","we%20are%20happy."},
		"null": {"", ""},
		"nospace": {"helloworld", "helloworld"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := replaceSpace(tt.source)
			if got != tt.want {
				t.Errorf("want=%s, got=%s\n", tt.want, got)
			}
		})
	}



}
