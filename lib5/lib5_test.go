package lib5

import "testing"

func Test_Sum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	if Sum(nums) != 15 {
		t.Errorf("sum of %v is 15, not %v", nums, Sum(nums))
	}
}
