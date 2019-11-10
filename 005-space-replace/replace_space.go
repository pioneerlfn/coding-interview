/*
@Time : 2019-11-10 20:31
@Author : lfn
@File : replace_space
*/

package _5_space_replace

func replaceSpace(arr string) string {
	if len(arr) == 0 {
		return ""
	}
	var spaceCount int
	for i := 0; i < len(arr); i++ {
		if arr[i] == ' ' {
			spaceCount++
		}
	}
	res := make([]byte, len(arr)+2*spaceCount)
	p1, p2 := len(arr)-1, len(res)-1
	for p1 != p2 {
		if arr[p1] != ' ' {
			res[p2] = arr[p1]
			p2--
		} else {
				res[p2] = '0'
				res[p2-1] = '2'
				res[p2-2] = '%'
				p2 = p2 - 3
		}
		p1--
	}
	copy(res[:p1+1], arr[:p1+1])
	return string(res)
}