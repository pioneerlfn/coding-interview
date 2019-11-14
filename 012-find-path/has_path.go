/*
@Time : 2019-11-14 22:08
@Author : lfn
@File : has_path
*/

package _12_find_path

func hasPath(matrix [][]byte, target string) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	if len(target) == 0 {
		return true // we treat this case as true
	}
	m, n := len(matrix), len(matrix[0])
	visited := make([]bool, m * n)
	var pos int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if hasPathCore(matrix, target, &pos, m, n, i, j, visited) {
				return true
			}
		}
	}
	return false
}

func hasPathCore(matrix [][]byte, target string, pos *int, rows, cols, row, col int, visited []bool) bool {
	// we have iterated all byte of the target, so has path.
	if *pos == len(target) {
		return true
	}
	var hasPath bool
	// if these four condition not satisfied, has no path.
	if row >= 0 && row < rows && col >= 0 && col < cols &&
		matrix[row][col] == target[*pos] && !visited[row*cols+col] {
		*pos++
		visited[row*cols+col] = true

		// may have path
		hasPath = hasPathCore(matrix, target, pos, rows, cols, row, col+1, visited) ||
			hasPathCore(matrix, target, pos, rows, cols, row, col-1, visited) ||
			hasPathCore(matrix, target, pos, rows, cols, row+1, col, visited) ||
			hasPathCore(matrix, target, pos, rows, cols, row-1, col, visited)
		if !hasPath {
			visited[row*cols+col] = false
			*pos--
		}
	}
	return hasPath
}

