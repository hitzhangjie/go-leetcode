package go_leetcode

func removeElement(nums []int, val int) int {

	if len(nums) == 0 {
		return 0
	}

	i := 0
	j := len(nums) - 1

	for ; i <= j; {

		if nums[i] == val {
			for j >= 0 && j < len(nums) && nums[j] == val {
				j--
			}

			if i == j+1 {
				nums = append(nums[0:i], nums[i+1:]...)
			} else {
				nums[i], nums[j] = nums[j], nums[i]
				j--
			}
		}

		i++
	}

	return j + 1
}
