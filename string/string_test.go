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
