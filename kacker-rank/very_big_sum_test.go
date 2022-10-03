package kacker_rank

import "testing"

func TestBigSum(t *testing.T) {
	t.Run("sum an array with some values", func(t *testing.T) {
		got := SumVeryBig([]int64{1, 2, 3})
		var want int64 = 6

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sum an array with no values", func(t *testing.T) {
		got := SumVeryBig([]int64{})
		var want int64 = 0

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sum an array with big values", func(t *testing.T) {
		got := SumVeryBig([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005})
		var want int64 = 5000000015

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
