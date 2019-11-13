/*
@Time : 2019-11-13 13:02
@Author : lfn
@File : stack_queue_test
*/

package _09_stack_queue

import (
	"math/rand"
	"reflect"
	"testing"
)


func TestStackQueue(t *testing.T) {
	tests := map[string]struct{
		in []byte
		out []byte
	}{
		"1":{[]byte{'a', 'b', 'c', 'd', 'e'},
			[]byte{'a', 'b', 'c', 'd', 'e'}},
		"2":{[]byte{'a', 'b', 'c', 'd', 'e', 'f'},
			[]byte{'a', 'b', 'c', 'd', 'e', 'f'}},
		"3":{[]byte{'a'},[]byte{'a'}},
		//"4":{[]byte{},[]byte{}},
	}
	var q Queue
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			n := rand.Intn(len(tt.in))
			var got []byte
			var i int
			for ; i < n; i++ {
				q.Push(tt.in[i])
			}
			if item := q.Dequeue(); item != '\x00' {
				got = append(got, item)
			}
			for ;i < len(tt.in); i++ {
				q.Push(tt.in[i])
			}
			item := q.Dequeue()
			for item != '\x00' {
				got = append(got, item)
				item = q.Dequeue()
			}
			if !reflect.DeepEqual(tt.out, got) {
				t.Errorf("want=%v, got=%v\n", tt.out, got)
			}
		})
	}
}










