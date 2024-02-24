package main

func main() {
	canJump([]int{2, 3, 1, 1, 4})
}
func canJump(nums []int) bool {
	maxDistant := 0
	for i := 0; i < len(nums); i++ {
		if i > maxDistant {
			return false
		}
		maxDistant = max(maxDistant, i+nums[i])
		if maxDistant >= len(nums)-1 {
			return true
		}
	}
	return false
}
