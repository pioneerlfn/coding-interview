/*
@Time : 2019-11-13 12:47
@Author : lfn
@File : stack_queue
*/

package _09_stack_queue

import (
	"fmt"
	"reflect"
)

type Stack []byte

func (s *Stack) Push(item byte) {
	*s = append(*s, item)
}

func (s *Stack) Pop() byte {
	n := len(*s)
	item := (*s)[n-1]
	*s = (*s)[:n-1]
	return item
}

type Queue struct {
	In Stack
	Out Stack
}

func (q *Queue) Push(item byte) {
	q.In = append(q.In, item)
}

func (q *Queue) Dequeue() byte {
	if len(q.Out) == 0 {
		for len(q.In) > 0 {
			q.Out.Push(q.In.Pop())
		}
	}
	if len(q.Out) > 0 {
		return q.Out.Pop()
	}
	return '\x00'
}

func main() {
	want := []byte{'a', 'b', 'c', 'd', 'e','f'}
	var got []byte
	var q Queue
	q.Push('a')
	q.Push('b')
	q.Push('c')
	got = append(got, q.Dequeue())
	got = append(got, q.Dequeue())
	q.Push('d')
	got = append(got, q.Dequeue())
	q.Push('e')
	q.Push('f')
	item := q.Dequeue()
	for item != '\x00' {
		got = append(got, item)
		item = q.Dequeue()
	}
	if reflect.DeepEqual(want, got) {
		fmt.Printf("want = %#v, got = %#v\n", want, got)
	}
}





