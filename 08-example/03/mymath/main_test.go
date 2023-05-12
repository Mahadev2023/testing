package mymath

import (
	"fmt"
	"testing"
)

// func TestCenteredAvg(t *testing.T) {
// 	a := []int{1, 4, 6, 8, 100}       //6
// 	b := []int{0, 8, 10, 1000}        //9
// 	c := []int{9000, 4, 10, 8, 6, 12} //9
// 	d := []int{123, 744, 140, 200}    //170
// 	e := [][]int{a, b, c, d}
// 	ans := []float64{6, 9, 9, 170}
// 	for i, v := range e {
// 		res := CenteredAvg(v)
// 		if res != ans[i] {
// 			t.Error("expected ", ans[i], "got", res)
// 		}
// 	}

// }

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		{[]int{1, 4, 6, 8, 100}, 6},
		{[]int{0, 8, 10, 1000}, 9},
		{[]int{9000, 4, 10, 8, 6, 12}, 9},
		{[]int{123, 744, 140, 200}, 170},
	}

	for _, ts := range tests {
		res := CenteredAvg(ts.data)
		if res != ts.answer {
			t.Error("expected ", ts.answer, "got", res)
		}
	}

}

func ExampleCenteredAvg() {
	a := []int{1, 4, 6, 8, 100}       //6
	b := []int{0, 8, 10, 1000}        //9
	c := []int{9000, 4, 10, 8, 6, 12} //9
	d := []int{123, 744, 140, 200}    //170
	xxi := [][]int{a, b, c, d}

	for _, v := range xxi {
		fmt.Println(CenteredAvg(v))
	}
	// Output
	// 6
	// 9
	// 9
	// 170

}

func BenchmarkCenteredAvg(b *testing.B) {
	a := []int{1, 4, 6, 8, 100}
	for i := 0; i < b.N; i++ {
		CenteredAvg(a)
	}
}
