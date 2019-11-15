/*
@Time : 2019-11-15 11:58
@Author : lfn
@File : count
*/

package _13_moving_count

func movingCount(threhold int, rows, cols int) int {
	if threhold == 0 || rows == 0 || cols == 0 {
		return 0
	}
	visited := make([]bool, rows*cols)
	count := movingCountCore(threhold, rows, cols, 0, 0, visited)
	return count
}

func movingCountCore(threhold int, rows, cols, row, col int, visited []bool) int {
	var count int
	if canVisit(threhold, rows, cols, row, col, visited) {
		visited[row*cols+col] = true
		count = 1 + movingCountCore(threhold, rows, cols, row, col+1, visited) +
			movingCountCore(threhold, rows, cols, row, col-1, visited) +
			movingCountCore(threhold, rows, cols, row+1, col, visited) +
			movingCountCore(threhold, rows, cols, row-1, col, visited)
	}
	return count
}

func canVisit(threhold int, rows, cols, row, col int, visited []bool) bool {
	if row >= 0 && row < rows && col >= 0 && col < cols &&
		validVisit(threhold, cols, row, col, visited) {
		return true
	}
	return false
}


func validVisit(threhold, cols, row, col int, visited []bool) bool {
	if visited[cols*row+col] {
		return false
	}
	rowSum := getDigitSum(row)
	colSum := getDigitSum(col)
	if rowSum + colSum >= threhold {
		return false
	}
	return true
}

func getDigitSum(num int) int {
	var sum int
	for num != 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}


