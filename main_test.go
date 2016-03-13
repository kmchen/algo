package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type test struct {
	input  []int
	output []int
}

func TestKsmallestNum(t *testing.T) {
	var testCases = []*test{
		&test{[]int{0}, []int{0}},
		&test{[]int{1, 0}, []int{1}},
		&test{[]int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48},
			[]int{4}},
		&test{[]int{3, 44, 38, 5, 47, 27, 2, 46, 4, 19, 40, 50, 48, 1},
			[]int{4}},
		&test{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{5}},
		&test{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{6}},
		&test{[]int{1, 3, 41, 24, 76, 11, 45, 3, 3, 64, 21, 69, 19, 36},
			[]int{21}},
	}
	ks := []int{1, 2, 3, 4, 5, 6, 7}
	Convey("Input should equal output", t, func() {
		for k, v := range testCases {
			input := v.input
			So(findKsmallestNum(input, ks[k]), ShouldResemble, v.output)
		}
	})
}
func TestFindPairsWithDiffK(t *testing.T) {
	var testCases = []*test{
		&test{[]int{3, 10, 20, 13, 7, 6, 7, 14, 16, 2, 4, 5, 11, 12},
			[]int{7}},
		&test{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{8}},
		&test{[]int{-10, -9, -8, -7, 0, 6, 5, 4, 3, 2, 1},
			[]int{9}},
	}
	ks := []int{3, 2, 1}
	Convey("Input should equal output", t, func() {
		for k, v := range testCases {
			input := v.input
			So(findPairsWithDiffK(input, ks[k]), ShouldResemble, v.output[0])
		}
	})
}
func TestConsecutiveInts(t *testing.T) {
	var testCases = []*test{
		&test{[]int{3, 10, 20, 13, 7, 6, 7, 14, 16, 2, 4, 5, 11, 12},
			[]int{7}},
		&test{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{8}},
		&test{[]int{-10, -9, -8, -7, 0, 6, 5, 4, 3, 2, 1},
			[]int{9}},
	}
	ks := []int{3, 2, 1}
	Convey("Input should equal output", t, func() {
		for k, v := range testCases {
			input := v.input
			So(findPairsWithDiffK(input, ks[k]), ShouldResemble, v.output[0])
		}
	})
}
func TestMaxSumSubArrays(t *testing.T) {
	var testCases = []*test{
		&test{[]int{1, -2, 3, 10, -4, 7, 2, -5},
			[]int{18}},
	}
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			input := v.input
			So(maxSumSubArrays(input), ShouldResemble, v.output[0])
		}
	})
}
func TestReverseWordsInSentence(t *testing.T) {
	var testCases = []struct {
		input, output string
	}{
		{"i am a student", "student a am i"},
		{"i", "i"},
		{"i ", "i"},
		{" i ", "i"},
		{"the sky is blue ", "blue is sky the"},
	}
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			input := v.input
			So(reverseWordsInSentence2(input), ShouldResemble, v.output)
		}
	})
}
func TestFirstCharOnlyOnce(t *testing.T) {
	var testCases = []struct {
		input, output string
	}{
		{"abaccdeff", "b"},
		{"a", "a"},
		{"abcdefgh", "a"},
		{"abcdefgabcdefgh", "h"},
	}
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			input := v.input
			So(firstCharOnlyOnce(input), ShouldResemble, v.output)
		}
	})
}
func TestLargestNumer(t *testing.T) {
	var testCases = []struct {
		input  []int
		output string
	}{
		{[]int{3, 30, 34, 5, 9}, "9534330"},
		{[]int{9, 30, 5, 3, 32}, "9533230"},
		{[]int{3, 30, 32, 5, 9}, "9533230"},
	}
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			input := v.input
			So(largestNumber(input), ShouldResemble, v.output)
		}
	})
}
func TestMergeSortedArray(t *testing.T) {
	//var testCases = []struct {
	//input  []int
	//output string
	//}{
	//{[]int{3, 30, 34, 5, 9}, "9534330"},
	//{[]int{9, 30, 5, 3, 32}, "9533230"},
	//{[]int{3, 30, 32, 5, 9}, "9533230"},
	//}
	//Convey("Input should equal output", t, func() {
	//for _, v := range testCases {
	//input := v.input
	//So(MergeSortedArray(input), ShouldResemble, v.output)
	//}
	//})
}
func TestMedianOfTwoSortedArray(t *testing.T) {
	var testCases = []struct {
		input1 []int
		input2 []int
		output int
	}{
		{[]int{1, 12, 15, 26, 38},
			[]int{2, 13, 17, 30, 45},
			16},
		{[]int{1, 2, 3, 6},
			[]int{4, 6, 8, 10},
			6},
	}
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			input1 := v.input1
			input2 := v.input2
			So(MedianOfTwoSortedArray(len(input1), input1, input2), ShouldResemble, v.output)
		}
	})
}
func TestTwoSum(t *testing.T) {
	var testCases = []struct {
		input  []int
		sum    int
		output []int
	}{
		{[]int{2, 11, 15, 7}, 9, []int{0, 3}},
		{[]int{0, 9, 15, 7}, 9, []int{0, 1}},
		{[]int{0, 9, 15, 7}, 9, []int{0, 1}},
		{[]int{2, 11, -2, 7}, 9, []int{1, 2}}}
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			input := v.input
			sum := v.sum
			So(TwoSum(input, sum), ShouldResemble, v.output)
		}
	})
}
