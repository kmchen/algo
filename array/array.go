package array

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
