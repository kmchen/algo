package main

import (
	"fmt"
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
			fmt.Println(input)
			So(quick(input), ShouldResemble, v.output)
		}
	})
}
