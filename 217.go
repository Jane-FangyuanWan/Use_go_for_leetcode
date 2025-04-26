package main

import "fmt"


func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums{
		if seen[num]{
			return true
		}
		seen[num] = true
	}
	return false
}

func main(){
	nums := []int{1, 2, 3, 4, 1}
	result := containsDuplicate(nums)
	fmt.Println(result)
}