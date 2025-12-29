package main

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		rowArr := make([]int, 0)
		for j := 0; j < cols; j++ {
			rowArr = append(rowArr, i*j)
		}
		matrix = append(matrix, rowArr)
	}
	return matrix
}

// SOLUTION WITHOUT APPEND(OPTIMIZED BUT LESS READABLE)

// func createMatrix(rows, cols int) [][]int {
// 	matrix := make([][]int, rows)
// 	for i := 0; i < rows; i++ {
// 		matrix[i] = make([]int, cols)
// 		for j := 0; j < cols; j++ {
// 			matrix[i][j] = i * j
// 		}
// 	}
// 	return matrix
// }
