package main

import "math"

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
