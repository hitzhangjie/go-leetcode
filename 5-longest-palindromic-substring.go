package go_leetcode

func longestPalindrome(s string) string {
	maxStr := ""
	maxLen := 0
	n := len(s)
	for i := 0; i <= n-1; i++ {
		for j := i; j <= n; j++ {
			xs := s[i:j]
			if isPalindrome(xs) && len(xs) > maxLen {
				maxLen = len(xs)
				maxStr = xs
			}
		}
	}
	return maxStr
}

func isPalindrome(s string) bool {
	n := len(s)

	if n == 0 {
		return false
	} else if n == 1 {
		return true
	}

	// 奇数
	if n%2 == 1 {
		mid := n / 2
		for i := mid - 1; i >= 0; i-- {
			if s[i] == s[2*mid-i] {
				continue
			}
			return false
		}
		return true
	}

	// 偶数
	mid := n/2 - 1
	for i := mid; i >= 0; i-- {
		if s[i] == s[n-i-1] {
			continue
		}
		return false
	}
	return true
}
