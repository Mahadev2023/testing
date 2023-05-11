package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	// tests := []test{
	// 	test{[]int{21, 21}, 42},
	// 	test{[]int{1, 2, 3}, 6},
	// 	test{[]int{0, 4, 0}, 4},
	// 	test{[]int{-1, -2, 3}, 0},
	// }

	tests := []test{
		{[]int{21, 21}, 42},
		{[]int{1, 2, 3}, 6},
		{[]int{0, 4, 0}, 4},
		{[]int{-1, -2, 3}, 0},
	}

	for _, v := range tests {

		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}

}
