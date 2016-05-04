package string

import (
	_ "../sort"
	"../util"
)

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

// Reverse String in O(n)
// Given a string reverse n first character
// Input : abcdef, 3
// output: defabc
// Solution 1: Loop n times and append to the end of the string
// Solution 2: // https://github.com/julycoding/The-Art-Of-Programming-By-July/blob/master/ebook/zh/01.01.md
func ReverseString1(str string, num int) string {
	var output string
	size := len(str)
	startIdx := size - num
	for i := startIdx; i < size; i++ {
		output += string(str[i])
	}
	for i := 0; i < num; i++ {
		output += string(str[i])
	}
	return output
}

// Solution 2 : Reverse [0:num-1] then reverse [num-1, 0]
func ReverseString2(str string, num int) string {
	left := ReverseString(str[0:num])
	right := ReverseString(str[num:])
	reversed := right + left
	diff := len(str) - num
	left = ReverseString(reversed[0:diff])
	right = ReverseString(reversed[diff:])
	return left + right
}

func ReverseString(str string) string {
	var output string
	idx := len(str) - 1
	for ; idx >= 0; idx-- {
		output += string(str[idx])
	}
	return output
}

// Implement StringContain(a, b string)
// Solution 1: Two for loops and look every single char of B in A in O(n*m)
func StringContainSol1(a, b string) bool {
	for _, v2 := range b {
		result := false
		for _, v1 := range a {
			if v2 == v1 {
				result = true
				break
			}
		}
		if !result {
			return false
		}
	}
	return true
}

// Solution 2: Sort both strings and repeat solution 1. Sorting takes O(nlgn)+O(mlgm)+O(n*m)
//func StringContainSol2(a, b string) bool {
//listA := strings.Split(a, "")
//listB := strings.Split(b, "")
//stringA := sort.Quick(listA)
//stringB := sort.Quick(listB)
//return StringContainSol1(strings.Join(stringA, ""), strings.Join(stringB, ""))
//}

// Solution 3: Use number theory. Prime numbers are multiplied for each character in string a, if the product can be divided by the prime number representation of characters in string B then string A contains string B
func StringContainSol3(a, b string) bool {
	primeNum := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239}
	prod := 1
	for _, v := range a {
		num := primeNum[v-'a']
		prod *= num
	}
	for _, v := range b {
		num := primeNum[v-'a']
		if prod%num != 0 {
			return false
		}
	}
	return true
}

// Solution 4: Use bitmap. Each character is represented as a single bit, we AND with bit representation of string B, if the result is 0 then string A does not contain string B.
func StringContainSol4(a, b string) bool {
	hash := 0
	for _, v := range a {
		hash |= 1 << uint(v-'a')
	}
	for _, v := range b {
		binaryRep := 1 << uint(v-'a')
		if hash&binaryRep == 0 {
			return false
		}
	}

	return true
}

// Implemnt atoi. Convert an integer in string to integer.
// Solution: Convert char by char and multiply by 10 for each subsequent character.
// Trick: Need to handle overflow. When given an extremely large number in string, the resulting integer might be overflowed. Need to handle +/- sign.
func Atoi(str string) int {
	// Validate +/- sign
	// Validate empty string
	if str == "" {
		return 0
	}
	// Assumption: int32
	MAX_INT := 1<<31 - 1
	S_MAX_INT := ^MAX_INT
	num := 0
	max := MAX_INT / 10
	unsigned := true
	if string(str[0]) == "-" {
		max = S_MAX_INT / 10
		unsigned = false
		str = str[1:]
	}
	for _, v := range str {
		// Validate non-integer
		if notDigit(v) {
			return 0
		}
		digit := runeToDigit(v)
		if unsigned {
			if num > max ||
				// When num = MAX_INT/10
				(num == max && MAX_INT%10 <= digit) {
				return MAX_INT
			}
			num = num*10 + digit
		} else {
			// When num < S_MAX_INT/10
			if num < max ||
				// When num = S_MAX_INT/10
				(num == max && -1*S_MAX_INT%10 <= digit) {
				return S_MAX_INT
			}
			num = num*10 - digit
		}
	}
	return num
}

func runeToDigit(num rune) int {
	return int(num - rune('0'))
}

func notDigit(num rune) bool {
	if num >= rune('0') && num <= rune('9') {
		return false
	}
	return true
}

// Solve palindrome, word spelles the same from backward. E.g. madam
// Solution1: Start from both end and check each char
// Solution2: Start from center towards both end
func palindromeSol1(str string) bool {
	rightPtr := len(str) - 1
	leftPtr := 0
	space := rune(' ')
	for rightPtr < len(str)/2 {
		right := rune(str[rightPtr])
		left := rune(str[leftPtr])
		// Handle white space
		if right == space {
			rightPtr++
			continue
		}
		if left == space {
			leftPtr--
			continue
		}
		if right != left {
			return false
		}
		leftPtr--
		rightPtr++
	}
	return true
}

func palindromeSol2(str string) bool {
	space := rune(' ')
	letterZ := rune('Z')
	lettera := rune('a')
	var newStr string
	for _, v := range str {
		if v >= lettera && v <= letterZ {
			newStr += string(v)
		}
	}

	var leftPtr int
	rightPtr := len(newStr) / 2
	if len(newStr)%2 == 0 {
		leftPtr = rightPtr - 1
	} else {
		leftPtr = rightPtr
	}
	for rightPtr < len(newStr) && leftPtr >= 0 {
		left := rune(newStr[leftPtr])
		right := rune(newStr[rightPtr])
		if left == space {
			leftPtr--
			continue
		}
		if right == space {
			rightPtr++
			continue
		}
		if left != right {
			return false
		}
		rightPtr++
		leftPtr--
	}
	return true
}

// TODO : Palindrome of link list
//分析：对于单链表结构，可以用两个指针从两端或者中间遍历并判断对应字符是否相等。但这里的关键就是如何朝两个方向遍历。由于单链表是单向的，所以要向两个方向遍历的话，可以采取经典的快慢指针的方法，即先位到链表的中间位置，再将链表的后半逆置，最后用两个指针同时从链表头部和中间开始同时遍历并比较即可。
// http://blog.csdn.net/u010305706/article/details/50810884

// Permutation
// Solution1: Fix one char at a time and swap each character and recursively print out all possible strings
// Solution2: Use next_permutation. Find
// https://github.com/julycoding/The-Art-Of-Programming-By-July/blob/master/ebook/zh/01.06.md
func LongestPermutation(str string) []string {
	return Permutation(str, 0)
}

func Permutation(str string, i int) []string {
	size := len(str)
	result := []string{}
	if i == size-1 {
		return []string{str}
	}
	for j := i + 1; j < size; j++ {
		result = append(result, Permutation(str, j)...)
		str = swap(str, i, j)
		result = append(result, Permutation(str, j)...)
		str = swap(str, j, i)
	}
	return result
}

func swap(str string, i, j int) string {
	if i == j || i < 0 || i > len(str)-1 {
		return str
	}
	var newStr string
	c := make(chan rune, 0)
	go func() {
		for k, v := range str {
			switch k {
			case i:
				c <- rune(str[j])
			case j:
				c <- rune(str[i])
			default:
				c <- v
			}
		}
	}()
	for _, _ = range str {
		newStr += string(<-c)
	}
	return newStr
}

// Postfix Expression Evaluation
// Input: 5 1 2 + 4 * + 3 -
// Output: 14
