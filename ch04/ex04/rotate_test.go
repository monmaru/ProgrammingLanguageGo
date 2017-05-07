package ex04

import "testing"

func TestEcho(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	want := []int{2, 3, 4, 5, 6, 1}

	rotate(s)

	for i, v := range s {
		if v != want[i] {
			t.Errorf("index = %d value = %d want = %d", i, v, want[i])
		}
	}
}
