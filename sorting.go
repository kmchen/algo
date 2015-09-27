package main

import "fmt"

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
