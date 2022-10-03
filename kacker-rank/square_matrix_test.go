package kacker_rank

import "testing"

func TestSquareMatrix(t *testing.T) {
	t.Run("sum main diagonal", func(t *testing.T) {
		matrix := [][]int32{
			{11, 2, 4},
			{4, 5, 6},
			{10, 8, -12},
		}
		got := DiagonalDifference(matrix)
		var want int32 = 15

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
