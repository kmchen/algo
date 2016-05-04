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
			So(palindromeSol1(input), ShouldResemble, v.output)
			So(palindromeSol2(input), ShouldResemble, v.output)
		}
	})
}

func TestLongestPalindrome(t *testing.T) {
	var testCases = []struct {
		input  string
		output []string
	}{
		{
			"12",
			[]string{"12", "21"},
		},
		{
			"123",
			[]string{"123", "213", "231", "321", "132", "312"},
		},
		{
			"54321",
			[]string{"12345", "12354", "12435", "12453", "12534", "12543", "13245", "13254", "13425", "13452",
				"13524", "13542", "14235", "14253", "14325", "14352", "14523", "14532", "15234", "15243",
				"15324", "15342", "15423", "15432", "21345", "21354", "21435", "21453", "21534", "21543",
				"23145", "23154", "23415", "23451", "23514", "23541", "24135", "24153", "24315", "24351",
				"24513", "24531", "25134", "25143", "25314", "25341", "25413", "25431", "31245", "31254",
				"31425", "31452", "31524", "31542", "32145", "32154", "32415", "32451", "32514", "32541",
				"34125", "34152", "34215", "34251", "34512", "34521", "35124", "35142", "35214", "35241",
				"35412", "35421", "41235", "41253", "41325", "41352", "41523", "41532", "42135", "42153",
				"42315", "42351", "42513", "42531", "43125", "43152", "43215", "43251", "43512", "43521",
				"45123", "45132", "45213", "45231", "45312", "45321", "51234", "51243", "51324", "51342",
				"51423", "51432", "52134", "52143", "52314", "52341", "52413", "52431", "53124", "53142",
				"53214", "53241", "53412", "53421", "54123", "54132", "54213", "54231", "54312", "54321",
			},
		},
	}
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			input := v.input
			check := func(output []string) func([]string) {
				return func(sliceStr []string) {
					for _, v := range sliceStr {
						So(output, ShouldContain, v)
					}
				}
			}
			check(v.output)(LongestPermutation(input))
		}
	})
}
