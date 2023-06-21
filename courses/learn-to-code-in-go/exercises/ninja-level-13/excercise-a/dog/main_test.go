package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	n := Years(10)
	cn := 10 * 7
	if n != cn {
		t.Error("got", n, "instead of", cn)
	}
}

func TestYearsTwo(t *testing.T) {
	n := Years(10)
	cn := 10 * 7
	if n != cn {
		t.Error("got", n, "instead of", cn)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
