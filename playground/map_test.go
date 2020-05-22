package playground

import "testing"


func TestCoordinate(t *testing.T) {
	c1 := Coordiante{1, 1}
	c2 := Coordiante{1, 2}
	c3 := Coordiante{1, 1}

	var result bool
	var want bool

	result = c1.equals(c2)
	want = false
	if result != want {
		t.Errorf("incorrect, got: %v, want: %v.", result, want)
	}

	result = c1.equals(c3)
	want = true
	if result != want {
		t.Errorf("incorrect, got: %v, want: %v.", result, want)
	}
}

