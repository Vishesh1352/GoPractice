package main

import "fmt"

var ret []int

func main() {
	nums := []int{3, 2, 4}
	target := 6
	ans := twoSum(nums, target)
	fmt.Println(ans)
}

func twoSum(nums []int, target int) []int {

	for i, v1 := range nums {
		if ret != nil {
			break
		}
		for j, v2 := range nums {
			if i != j {
				sum := v1 + v2
				if sum == target {

					ans := []int{i, j}

					ret = append(ret, ans[0], ans[1])

					break
				}
			}
		}
	}
	return ret
}
