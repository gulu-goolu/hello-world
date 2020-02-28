package lib4

import "testing"

func Test_Div(t *testing.T) {
	d := Div(6, 2)
	if d != 3 {
		t.Errorf("6 / 2 should equal to 3, not %v", d)
	}
}
