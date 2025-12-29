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
