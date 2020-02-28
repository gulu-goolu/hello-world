package lib2

import "testing"

func TestSub(t *testing.T) {
	if 2 != Sub(5, 3) {
		t.Errorf("5 - 3 should be 2")
	}
}
