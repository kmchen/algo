package string

import util "../util"

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

// [LeetCode 5] Longest Palindromic Substring
// Given a string S, find the longest palindromic substring in S.
// You may assume that the maximum length of S is 1000, and there
// exists one unique longest palindromic substring.
// Solution 1: Use DP, build a 2D array. Rows are left hand side and columns are right hand side and fill up the upper part of the array
// SOlution 2: Search around corner is basically iterate through the entire string and assume each char (or char pair) as center of the palindrome.
// Input : BBABCBCAB, then the output should be 7 as “BABCBAB"
func LongestPalindromicSubstring(str string) int {
	size := len(str)
	// Make 2d slices
	matrix := [][]int{}
	// Set diagonal cells to 1 since the longest palindromic
	// substring of length 1 is in fact 1
	for k, _ := range str {
		matrix = append(matrix, make([]int, size))
		matrix[k][k] = 1
	}
	// Length of substring
	for i := 1; i < size; i++ {
		// Loop through each char as starting index
		for start, _ := range str {
			end := start + i
			if end >= size {
				continue
			}
			// if starting char equals to end char
			// current cell = 2 + inner substring
			// e.g. a(bb)a
			if str[start] == str[end] {
				matrix[start][end] = 2 + matrix[start+1][end-1]
			} else {
				// if starting char and end char are not equal
				// current cell = the max of substrings below
				// e.g. Max of (abb)b and a(bbb)
				matrix[start][end] = util.Max(matrix[start][end-1], matrix[start+1][end])
			}
		}
	}
	return matrix[0][size-1]
}
