package main

import (
	"fmt"
	"sort"
)

func WhichIndex(nums []int, target int) []int {
	result := make(map[int]int)

	for i, num := range nums {
		if val, ok := result[target-num]; ok {
			return []int{val, i}
		}
		result[num] = i
	}

	return nil
}

func Zero(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func main() {
	result1 := WhichIndex([]int{2, 7, 11, 15}, 9)
	fmt.Println(result1)

	result2 := Zero([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(result2)

	// maaf untuk nomor 3 saya tidak mengerti maksudnya
}
