package sum

import "testing"

func TestMultipleNumbers(t *testing.T) {
	type test struct {
		data   []int32
		answer int32
	}

	tests := []test{
		test{[]int32{1, 2, 3}, 6},
		test{[]int32{-1, -2}, -3},
		test{[]int32{0, 5, -1, 3, 5}, 12},
	}

	for _, v := range tests {
		x, err := Calculate(v.data)
		if err != nil {
			t.Error(err)
		}
		if x != v.answer {
			t.Error("Expected:", v.answer, "Got:", x)
		}
	}
}

func TestEmptyArray(t *testing.T) {
	data := []int32{}
	var answer int32
	answer = 0
	x, err := Calculate(data)
	if err != nil {
		t.Error(err)
	}
	if x != answer {
		t.Error("Expected:", answer, "Got:", x)
	}
}
