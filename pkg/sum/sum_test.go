package sum

import "testing"

func TestMultipleNumbers(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{1, 2, 3}, 6},
		test{[]int{-1, -2}, -3},
		test{[]int{0, 5, -1, 3, 5}, 12},
	}

	for _, v := range tests {
		x, err := Sum(v.data)
		if err != nil {
			t.Error(err)
		}
		if x != v.answer {
			t.Error("Expected:", v.answer, "Got:", x)
		}
	}
}

func TestEmptyArray(t *testing.T) {
	data := []int{}
	answer := 0
	x, err := Sum(data)
	if err != nil {
		t.Error(err)
	}
	if x != answer {
		t.Error("Expected:", answer, "Got:", x)
	}
}
