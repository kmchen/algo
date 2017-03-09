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
