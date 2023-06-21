package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		test{
			data:   []int{10, 20, 40, 60, 80},
			answer: 40,
		},
		test{
			data:   []int{2, 4, 6, 8, 10, 12},
			answer: 7,
		},
		test{
			data:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			answer: 5,
		},
	}

	for _, v := range tests {
		f := CenteredAvg(v.data)
		if f != v.answer {
			t.Error("got", f, "wanted", v.answer)
		}
	}

}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 4, 6, 9}))
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10000})
	}
}
