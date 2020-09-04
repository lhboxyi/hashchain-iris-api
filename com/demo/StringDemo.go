package demo

/**
无重复字符的最长子串
 */
func LengthOfLongestSubstring(s string) int {
	var count = make(map[rune]int) // 存放字符的位置
	var max int
	left := 0
	for k, v := range s {
		if val, ok := count[v]; ok {
			if left < val {
				left = val
			}
		}
		if max < k-left+1 {
			max = k - left + 1
		}
		count[v] = k + 1 // 更新位置
	}
	return max
}
