package go_leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 假定nums1长度小于等于nums2
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	// n1, n2, n3, | n4, n5, n6
	//             -----
	// m1, m2, m3, m4, | m5, m6, m7, m8

	// 因为nums1, nums2有序，只需要找到两个分界线 midA, midB，保证
	// max(nums1[midA - 1], nums2[midB - 1]) < min(nums1[midA], nums2[midB])

	// 因为总数是可以算出来的，那么分界线左边的元素数量也是可以算出来的：
	// - 如果总数为奇数个的话，那么分界线左侧可以比右边多一个，其中最大的那个就是中位数
	// - 如果总数为偶数个的话，那么分界线左侧和右侧数量一样多，取左边最大的、右边最小的两个的平均值
	m, n := len(nums1), len(nums2)
	totalLeft := (m + n + 1) / 2

	// 我们固定先扫描短的数组nums1，直接判断分界线两侧的数是否满足要求即可
	left, right := 0, m

	for left < right {
		//println(left, right)
		i := left + (right-left+1)/2 // 表示nums1分界线，其实也就是左侧元素数量
		j := totalLeft - i           // 表示nums2分界线，因为总数是固定的，所以可以直接确定其值

		//println(i, j)
		if nums1[i-1] > nums2[j] {
			right = i - 1
		} else {
			left = i
		}
	}
	// 更新right（left始终是在更新的，right不是）
	right = totalLeft - left

	var (
		nums1LeftMax  int
		nums1RightMin int
		nums2LeftMax  int
		nums2RightMin int
	)

	if left == 0 {
		nums1LeftMax = MinInt64
	} else {
		nums1LeftMax = nums1[left-1]
	}

	if left == m {
		nums1RightMin = MaxInt64
	} else {
		nums1RightMin = nums1[left]
	}

	if right == 0 {
		nums2LeftMax = MinInt64
	} else {
		nums2LeftMax = nums2[right-1]
	}

	if right == n {
		nums2RightMin = MaxInt64
	} else {
		nums2RightMin = nums2[right]
	}

	//println(nums1LeftMax, nums1RightMin)
	//println(nums2LeftMax, nums2RightMin)

	if (m+n)&1 == 0 {
		// 偶数个，返回分界线左边最大数和右边最小数的平均值
		return ((float64)(max(nums1LeftMax, nums2LeftMax)) + ((float64)(min(nums1RightMin, nums2RightMin)))) / 2
	} else {
		// 奇数个，返回分界线左边最大数即可
		return float64(max(nums1LeftMax, nums2LeftMax))
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

const (
	MaxInt64 = 1<<63 - 1
	MinInt64 = -1 << 63
)
