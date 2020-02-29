package lib5

func Sum(nums []int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}
