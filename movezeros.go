package go_algo

func MoveZeros(nums []int) []int {
	l1 := len(nums)
	if l1 == 0 {
		return nums
	}

	to := l1
	for i := 0; i < to; {
		if nums[i] == 0 {
			nums = append(nums[0:i], nums[i+1:l1]...)
			nums = append(nums, 0)
			to--
		} else {
			i++
		}
	}

	return nums
}
