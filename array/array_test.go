package array

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindPairWithGivenSum(t *testing.T) {
	type test struct {
		input  []int
		sum    int
		output []int
	}
	testCases := []*test{
		//&test{[]int{}, []int{}},
		//&test{[]int{0}, []int{0}},
		//&test{[]int{1, 0}, []int{0, 1}},
		&test{
			[]int{8, 7, 2, 5, 3, 1},
			10,
			[]int{0, 2, 1, 4, 3, 3},
		},
	}
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			result := FindPairWithGivenSum(v.input, v.sum)
			for _, val := range result {
				So(val, ShouldBeIn, v.output)
			}
		}
	})
}

func TestSubArrayWithZeroSum(t *testing.T) {
	type test struct {
		input  []int
		output [][]int
	}
	//testCases := []*test{
	//&test{
	//[]int{4, 2, -3, -1, 0, 4},
	//[][]int{
	//[]int{-3, -1, 0, 4},
	//[]int{0},
	//},
	//},
	//&test{
	//[]int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2},
	//[][]int{
	//[]int{3, 4, -7},
	//[]int{4, -7, 3},
	//[]int{-7, 3, 1, 3},
	//[]int{3, 1, -4},
	//[]int{3, 1, 3, 1, -4, -2, -2},
	//[]int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2},
	//},
	//},
	//}
	input := []int{4, 2, -3, -1, 0, 4}
	SubArrayWithZeroSum(input)
	input = []int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2}
	SubArrayWithZeroSum(input)
	//Convey("Input should equal output", t, func() {
	//for _, v := range testCases {
	//result := SubArrayWithZeroSum(v.input)
	//for _, val := range result {
	//So(val, ShouldResemble, v.output)
	//}
	//}
	//})
}

//Find a duplicate element in a limited range array
func testFindDuplicates(t *testing.T) {
	type test struct {
		input  []int
		output []int
	}
	testCases := []*test{
		&test{
			[]int{1, 2, 3, 4, 4},
			[]int{4},
		},
		&test{
			[]int{1, 2, 3, 4, 2},
			[]int{2},
		},
	}
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			result := FindDuplicates(v.input)
			So(result, ShouldEqual, v.output)
		}
	})
}
