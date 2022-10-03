package kacker_rank

func DiagonalDifference(matrix [][]int32) int32 {

	var total = getSumOfDiagonals(matrix)
	if total < 0 {
		total = total * -1
	}
	return total
}

func getSumOfDiagonals(matrix [][]int32) int32 {
	var sumMainDiagonal int32 = 0
	var sumInverseDiagonal int32 = 0
	matrixLen := len(matrix)

	for i := 0; i < len(matrix); i++ {
		sumMainDiagonal += matrix[i][i]
		matrixLen--
		sumInverseDiagonal += matrix[i][matrixLen]
	}

	var total = sumMainDiagonal - sumInverseDiagonal
	return total
}
