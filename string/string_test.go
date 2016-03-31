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
