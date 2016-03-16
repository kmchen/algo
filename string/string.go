package string

// [LeetCode 3] Longest Substring Without Repeating Characters
// Given a string, find the length of the longest substring without
// repeating characters. For example, the longest substring without
// repeating letters for “abcabcbb” is “abc”, which the length is 3.
// For “bbbbb” the longest substring is “b”, with the length of 1.
// Input : “bbbbb” => length = 1
// Input : “abcabcbb” => length = 3
// Input : “aabcabcbb” => length = 3
// Solution : Sliding window
func LongestSubstring(input string) int {
	// String length >= 2
	left, right := 0, 1
	longest := 0
	size := len(input)
	for ; right < size; right++ {
		// left != right
		if input[left] != input[right] {
			longest = right - left
			if left == 0 {
				longest++
			}
			continue
		}
		// left == right
		left++
	}
	if longest == 0 {
		return 1
	}
	return longest
}
