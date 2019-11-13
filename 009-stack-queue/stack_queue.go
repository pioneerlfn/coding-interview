/*
@Time : 2019-11-13 12:47
@Author : lfn
@File : stack_queue
*/

package _09_stack_queue

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
