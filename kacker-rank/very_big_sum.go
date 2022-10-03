package kacker_rank

func SumVeryBig(array []int64) int64 {

	var count int64 = 0
	for i := 0; i < len(array); i++ {
		count += array[i]
	}
	return count
}
