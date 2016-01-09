// Problems from http://codercareer.blogspot.tw/2011/10/no-13-first-character-appearing-only.html
package main

import (
	"fmt"
	"math"
	"strings"
)

func bubble(array []int) []int {
	if len(array) == 0 || len(array) == 1 {
		return array
	}
	endIdx := len(array) - 1
	for {
		i := 0
		if endIdx <= i {
			break
		}
		for ; i < endIdx; i++ {
			if array[i] >= array[i+1] {
				tmp := array[i]
				array[i] = array[i+1]
				array[i+1] = tmp
			}
		}
		endIdx = endIdx - 1
	}
	return array
}

func selectSort(array []int) []int {
	if len(array) == 1 || len(array) == 0 {
		return array
	}
	end := len(array) - 1
	for j := 0; j < end; j++ {
		min := array[j]
		minIdx := j
		for i := j; i < len(array); i++ {
			if array[i] < min {
				min = array[i]
				minIdx = i
			}
		}
		tmp := array[j]
		array[j] = array[minIdx]
		array[minIdx] = tmp
	}
	return array
}

func mergeSort(array []int) []int {
	if len(array) == 1 || len(array) == 0 {
		return array
	}
	middle := len(array) / 2
	low := 0
	high := len(array)
	right := mergeSort(array[low:middle])
	left := mergeSort(array[middle:high])
	return merge(right, left)
}

func merge(a ...[]int) []int {
	if len(a) != 2 {
		return nil
	}
	right := a[0]
	left := a[1]
	var merged = make([]int, len(right)+len(left))
	var r, l = 0, 0
	for i := 0; i < len(right)+len(left); i++ {
		if r >= len(right) {
			merged[i] = left[l]
			l++
			continue
		}
		if l >= len(left) {
			merged[i] = right[r]
			r++
			continue
		}
		if right[r] > left[l] {
			merged[i] = left[l]
			l++
		} else {
			merged[i] = right[r]
			r++
		}
	}
	return merged
}

func quick(a []int) []int {
	if len(a) == 0 || len(a) == 1 {
		return a
	}
	start, end := 0, len(a)-1
	if start == end-1 {
		if a[start] > a[end] {
			swap(a, start, end)
			return a
		}
	}
	start, pivot, end := sort(a, start, end)
	quick(a[start : pivot+1])
	quick(a[pivot+1 : end+1])
	return a
}

func sort(a []int, pivot, end int) (int, int, int) {
	count := pivot + 1
	for j := count; j <= end; j++ {
		if a[j] < a[pivot] {
			swap(a, count, j)
			count++
		}
	}
	start := pivot
	pivot = count - 1
	swap(a, start, pivot)
	return start, pivot, end
}

func swap(a []int, i, j int) {
	if i == j {
		return
	}
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

// Find ‘k’ smallest numbers from a million numbers
// Solution 1 : Sort in O(nlogn)
// Solution 2 : Heapify to max or min heap
// Solution 3 : Use quick sort to find partition takes O(n)
func findKsmallestNum(slice []int, k int) []int {
	if len(slice) == 0 {
		return nil
	}
	kth := k - 1
	pivot := 0
	start, end := 1, len(slice)-1
	// comparison
	for {
		// Scan from right to left
		for slice[end] > slice[pivot] {
			end--
		}
		// Scan from left to right
		for start <= len(slice)-1 && slice[start] <= slice[pivot] {
			start++
		}
		if start >= end {
			break
		}
		slice[start], slice[end] = slice[end], slice[start]
	}
	// partition
	slice[end], slice[pivot] = slice[pivot], slice[end]
	if end < kth {
		start = end + 1
		return findKsmallestNum(slice[start:], k-len(slice[:start]))
	} else if end > kth {
		return findKsmallestNum(slice[:end+1], k)
	}
	return []int{slice[end]}
}

// You have been given N numbers and a value K.
// Find the number of pairs in this given list of numbers whose difference is exactly equal to K.
// Solution 1: Brute force. Loop through array and find the element with distance K
// Solution 2: Sort the array. Use binary search to find the pair
// Solution 3: Sort the array. Start from the last element a[N], find the value J at index I.
// 			   For element a[N-1], find the value X in the range a[0:J-1]
func findPairsWithDiffK(slice []int, dist int) int {
	if len(slice) <= 1 {
		return 0
	}
	var count int
	sorted := quick(slice)
	size := len(sorted) - 1
	index := -1
	for i := size; i > 0; i-- {
		ele := sorted[i]
		var inner []int
		if index == -1 {
			index = len(sorted)
		}
		inner = sorted[:index]
		innerLen := len(inner) - 1
		for j := innerLen; j >= 0; j-- {
			if j != innerLen && inner[j] == inner[j+1] {
				continue
			}
			if abs(inner[j]) < abs(ele)-dist {
				break
			}
			if abs(abs(inner[j])-abs(ele)) == dist {
				index = j
				count++
			}
		}
	}
	return count
}

func abs(num int) int {
	return int(math.Abs(float64(num)))
}

type Tree interface {
	Insert(int, *Node)
	Remove(*Node)
}
type TreeImpl struct {
	root    *Node
	current *Node
	count   int
	depth   int
}
type Node struct {
	left, right *Node
	value       int
}

func (t *TreeImpl) Insert(idx int, n *Node) {
	if t.count == 0 && idx == 0 {
		t.count++
		t.root = n
		return
	}
	// Add depth when total num of nodes equals 2^(d+1) - 1
	if t.count == pow(2, t.depth+1)-1 {
		t.depth++
	}
	t.count++

	// Construct path to node
	path := make([]int, 0)
	var parent = t.root
	var current = t.root
	constructPath(idx, path)

	for k := len(path) - 1; k > 0; k-- {
		parent = current
		if k == 1 {
			break
		}
		if k%2 == 0 {
			current = parent.left
		} else {
			current = parent.right
		}
	}

	// Insert to left or right
	if path[0]%2 == 0 {
		parent.left = n
	} else {
		parent.right = n
	}
	return
}

func constructPath(idx int, slice []int) {
	if idx <= 0 {
		return
	}
	var parent int
	if idx%2 == 0 {
		slice = append(slice, 0)
		parent = idx / 2
	} else {
		slice = append(slice, 1)
		parent = (idx - 1) / 2
	}
	constructPath(parent, slice)
	return
}

func pow(b int, p int) int {
	return int(math.Pow(float64(b), float64(p))) - 1
}

func (t *TreeImpl) Remove(n *Node) {
	return
}

// Fact : a complete binary tree of depth d equals 2^(d+1) – 1 nodes
// Since all leaves in such a tree are at level d,
// the tree contains 2^d leaves and, therefore, 2^d - 1 internal nodes
func binaryHeap(slice []int) Tree {
	if len(slice) <= 0 {
		return nil
	}
	heap := new(TreeImpl)
	for k, v := range slice {
		n := &Node{nil, nil, v}
		heap.Insert(k, n)
	}
	return heap
}

// Check if Array is consecutive Integers
// 21,24,22,26,23,25 -> True
// Solution 1 : Soriting n O(nlogn)
// Solution 2 : Store in another array and all values can't be > 2
// 				Time complexity : O(n), Space complexity : O(n)
// Solution 3 : Find the min. For each value in the array, substract by min
// 				array[array[i]] = -1, if array[array[i]] is already negative
// 				then the array is not consecutive
func consecutiveInt(slice []int) bool {
	var min, max = 0, 0

	// Find min O(n)
	for k, v := range slice {
		if k == 0 {
			min = v
			max = v
			continue
		}
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	// TODO : Check if max - min +1 == array length
	if max-min+1 != len(slice) {
		return false
	}

	// Substract by min O(n)
	for k, v := range slice {
		slice[k] = v - min
	}
	// Mark array[i] to -1, and detect duplicates
	for k := range slice {
		if slice[slice[k]] == -1 {
			return false
		}
		slice[slice[k]] = -1
	}
	return true
}

// Maximum Sum of All Sub-arrays
// For example, in the array {1, -2, 3, 10, -4, 7, 2, -5},
// its sub-array {3, 10, -4, 7, 2} has the maximum sum 18.
// Solution 1:  Enumerating all of sub-arrays and calculate their sum
// An array with n elements has n(n+1)/2 sub-arrays. It costs O(n2) time at least to calculate their sum.
// Solution 2: Analyse array number by number. Sum is equal to the accumulated sum, if current sum < accumulated sum
func maxSumSubArrays(slice []int) int {
	maxSum, sum := 0, 0
	for k, v := range slice {
		fmt.Println(k, v, maxSum, sum)
		if k == 0 {
			maxSum, sum = v, v
			continue
		}
		var tmp = sum + v
		if tmp < sum {
			sum += v
			continue
		}
		sum += v
		if maxSum < sum && sum < v {
			maxSum = v
			sum = v
			continue
		}
		maxSum = sum
	}
	return maxSum
}

// String : Reverse words in a sentence
// Input	: i am a student
// Outout : student a am i
func reverseWordsInSentence(input string) string {
	size := len(input) - 1
	var revString, result, tmp string
	for i := range input {
		index := size - i
		char := string(input[index])
		if i == size {
			result += char
			break
		}
		if char != " " {
			tmp += char
			continue
		}
		// Got reversed string e.g. tneduts
		for _, v := range tmp {
			revString = string(v) + revString
		}
		result += revString + " "
		tmp = ""
		revString = ""
	}
	return result
}

// String : Reverse words in a sentence
// Input	: "the sky is blue"
// Outout : "blue is sky the"
func reverseWordsInSentence2(input string) string {
	result := ""
	head := len(input) - 1
	tail := len(input)
	// Loop through the sentence
	for idx := len(input) - 1; idx >= 0; idx-- {
		// Record index when hits an empty space
		if string(input[idx]) != " " {
			head--
			continue
		}
		// Append substring[head:tail] to the result
		// Reset tail = head
		result += input[head:tail]
		tail = head
		head--
	}
	result += " " + input[0:tail]
	return strings.Trim(result, " ")
}

// String : First character appears only once
// Problem :Implement a function to find the first character in a string which only appears once.
// Input : "abaccdeff"
// Output : "b"
func firstCharOnlyOnce(input string) string {
	result := make(map[rune]int)
	var char string
	for _, v := range input {
		result[v]++
	}
	for _, v := range input {
		if result[v] == 1 {
			char = string(v)
			break
		}
	}
	return char
}

//Problem: Left rotation of a string is to move some leading characters to its tail. Please implement a function to rotate a string.
//For example, if the input string is “abcdefg” and a number 2, the rotated result is “cdefgab”.
