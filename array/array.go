package array

import "fmt"

// Find pair with given sum in the array
// Given an unsorted array of integers, find a pair with given sum in it
// Input: [8, 7, 2, 5, 3, 1], sum = 10
// Output: (0, 2), (1, 5)
// O(n^2): Consider every pair in given array results
// O(nlogn): Sort first then having a low and high index calculating the newSum
// Increment low when the sum>newSum and decrement high when newSum<sum. Stop when low > high
// O(n): Hashing
func FindPairWithGivenSum(input []int, sum int) []int {
	result := make([]int, 0)
	// O(nlogn) solution
	//idxMap := map[int]int{}
	//for key, val := range input {
	//idxMap[val] = key
	//}
	//sortedInput := sort.Quick(input)
	//low, high := 0, len(input)-1
	//for low <= high {
	//newSum := sortedInput[low] + sortedInput[high]
	//if newSum == sum {
	//result = append(result, idxMap[sortedInput[low]], idxMap[sortedInput[high]])
	//}
	//if newSum > sum {
	//high--
	//} else {
	//low++
	//}
	//}

	// O(n) hashing
	hash := make(map[int]int)
	for key, val := range input {
		diff := sum - val
		if idx, ok := hash[diff]; ok {
			result = append(result, key, idx)
		}
		hash[val] = key
	}
	return result
}

// Find sub-array with 0 sum
// Given an array of integers, check if array contains a sub-array having 0 sum. Also, prints end-points of all such sub-arrays.
// O(n^2): For each element, loop through all elements in subarray
// O(n): Traverse the given array, and maintain sum of elements seen so far in a map.
// If sum is seen before, there exists at-least one sub-array with 0 sum which ends at current index and we print all such sub-arrays
func SubArrayWithZeroSum(input []int) {
	// O(n) multi-dimensional map
	set := map[int][]int{0: []int{-1}}
	sum := 0
	for key, val := range input {
		sum += val
		if indexes, ok := set[sum]; ok {
			for _, idx := range indexes {
				fmt.Println(idx+1, " ... ", key)
			}
		}
		set[sum] = append(set[sum], key)
	}
}

// Find a duplicate element in a limited range array
// Given a limited range array of size n where array contains elements between 1 to n-1 with one element repeating
// find the duplicate number in it.
func FindDuplicates(input []int) int {
	// using XOR
	result := 0
	//for _, val := range input {
	//result ^= val
	//}
	//for key, _ := range input {
	//result ^= key
	//}

	// using constant space
	for _, val := range input {
		if input[val] < 0 {
			return input[val]
		}
		input[val] = input[val] * -1
	}
	return result
}

// Find largest sub-array formed by consecutive integers
// Given an array of integers, find largest sub-array formed by consecutive integers. The sub-array should contain all distinct values.
