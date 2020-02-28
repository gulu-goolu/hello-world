package lib3

import "testing"

func TestMul(t *testing.T) {
	if Mul(2, 3) != 6 {
		t.Errorf("2 * 3 should equal to 6")
	}
}
