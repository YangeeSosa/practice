package kata

func MinMax(arr []int) [2]int {
	max := arr[0]
	min := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		} else if arr[i] > max {
			max = arr[i]
		}
	}
	return [2]int{min, max}
}
