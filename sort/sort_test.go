package sort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type test struct {
	input  []int
	output []int
}

var testCases = []*test{
	&test{[]int{}, []int{}},
	&test{[]int{0}, []int{0}},
	&test{[]int{1, 0}, []int{0, 1}},
	&test{[]int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48},
		[]int{2, 3, 4, 5, 15, 19, 26, 27, 36, 38, 44, 46, 47, 48, 50}},
	&test{[]int{3, 44, 38, 5, 47, 27, 2, 46, 4, 19, 40, 50, 48, 1},
		[]int{1, 2, 3, 4, 5, 19, 27, 38, 40, 44, 46, 47, 48, 50}},
	&test{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	&test{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	&test{[]int{1, 3, 41, 24, 76, 11, 45, 3, 3, 64, 21, 69, 19, 36},
		[]int{1, 3, 3, 3, 11, 19, 21, 24, 36, 41, 45, 64, 69, 76}},
}

func testBubble(t *testing.T) {
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			input := v.input
			So(bubble(input), ShouldResemble, v.output)
		}
	})
}
func TestSelect(t *testing.T) {
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			input := v.input
			So(selectSort(input), ShouldResemble, v.output)
		}
	})
}
func TestMerge(t *testing.T) {
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			input := v.input
			So(mergeSort(input), ShouldResemble, v.output)
		}
	})
}
func TestQuick(t *testing.T) {
	Convey("Input should equal output", t, func() {
		for _, v := range testCases {
			input := v.input
			So(quick(input), ShouldResemble, v.output)
		}
	})
}
