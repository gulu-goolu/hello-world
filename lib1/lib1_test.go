package lib1

import "testing"

func TestAdd(t *testing.T) {
	if 5 != Add(2, 3) {
		t.Errorf("%v, %v not equal", 2, 3)
	}
}
