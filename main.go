package main

import "fmt"

func oddSumMaxPair(nums []int) []int {
	var maxSum, pairSum int
	var result []int

	for i:=0; i<len(nums); i++ {
		for j:=i+1; j<len(nums); j++ {
			pairSum = nums[i] + nums[j]
			if pairSum % 2 == 1 && pairSum > maxSum {
				maxSum = pairSum
				result = []int{nums[i],nums[j]}
			}
		}
	}

	fmt.Println(result)
	return result
}

func main() {
	oddSumMaxPair([]int{7, 3, 9, 7, 5, 1, 5})
}