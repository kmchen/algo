package string

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLongestSubstring(t *testing.T) {

	// Input : “bbbbb” => length = 1
	// Input : “abcabcbb” => length = 3
	// Input : “aabcabcbb” => length = 3
	var testCases = []struct {
		input  string
		output int
	}{
		{"bbbbbb", 1},
		{"abcabcb", 3},
		{"aabcabcb", 3},
	}
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			input := v.input
			So(LongestSubstring(input), ShouldResemble, v.output)
		}
	})
}
func TestLongestPalindromicSubstring(t *testing.T) {
	// Input : BBABCBCAB, then the output should be 7 as “BABCBAB"
	var testCases = []struct {
		input  string
		output int
	}{
		{
			"BBABCBCAB",
			7, //"BABCBAB"
		},
		{
			"abac",
			3,
		},
		{
			"caba",
			3,
		},
		{
			"agbdba",
			5,
		},
	}
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			input := v.input
			So(LongestPalindromicSubstring(input), ShouldResemble, v.output)
		}
	})
}
func TestReverseString(t *testing.T) {
	var testCases = []struct {
		input  string
		n      int
		output string
	}{
		{
			"abcdef",
			3,
			"defabc",
		},
		{
			"abcdef",
			6,
			"abcdef",
		},
		{
			"abcdef",
			0,
			"abcdef",
		},
	}
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			input := v.input
			So(ReverseString2(input, v.n), ShouldResemble, v.output)
		}
	})
}
func TestStringContain(t *testing.T) {
	var testCases = []struct {
		input1 string
		input2 string
		output bool
	}{
		{
			"abcdef",
			"bad",
			true,
		},
		{
			"abcd",
			"bce",
			false,
		},
	}
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			input1 := v.input1
			input2 := v.input2
			So(StringContainSol1(input1, input2), ShouldResemble, v.output)
			So(StringContainSol3(input1, input2), ShouldResemble, v.output)
			So(StringContainSol4(input1, input2), ShouldResemble, v.output)
		}
	})
}
func TestAtoi(t *testing.T) {
	var testCases = []struct {
		input  string
		output int
	}{
		{
			"2147483647",
			2147483647,
		},
		{
			"2147483648",
			2147483647,
		},
		{
			"2147483648000",
			2147483647,
		},
		{
			"2147483646",
			2147483646,
		},
		{
			"2147483",
			2147483,
		},
		{
			"-2147483648",
			-2147483648,
		},
		{
			"-2147483647",
			-2147483647,
		},
		{
			"-2147483649",
			-2147483648,
		},
		{
			"-214748364900000",
			-2147483648,
		},
		{
			"-2147483",
			-2147483,
		},
	}
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			input := v.input
			So(Atoi(input), ShouldResemble, v.output)
		}
	})
}

func TestPalindrome(t *testing.T) {
	var testCases = []struct {
		input  string
		output bool
	}{
		{
			"A man, a plan, a canal, Panama!",
			true,
		},
		{
			"Amor, Roma",
			true,
		},
		{
			"race car",
			true,
		},
		{
			"stack cats",
			true,
		},
		{
			"step on no pets",
			true,
		},
		{
			"taco cat",
			true,
		},
		{
			"put it up",
			true,
		},
	}
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			input := v.input
			So(palindrome(input), ShouldResemble, v.output)
		}
	})
}
