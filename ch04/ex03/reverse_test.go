package ex03

import "testing"

func TestEcho(t *testing.T) {
	s := [6]int{1, 2, 3, 4, 5, 6}
	want := [6]int{6, 5, 4, 3, 2, 1}

	reverse(&s)

	for i, v := range s {
		if v != want[i] {
			t.Errorf("index = %d value = %d want = %d", i, v, want[i])
		}
	}
}
