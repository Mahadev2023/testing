package dog

import (
	"fmt"
	"testing"
)

func TestYear(t *testing.T) {
	n := Years(5)
	if n != 35 {
		t.Error("expected", 35, "got", 5)
	}
}

func TestYearsTwo(t *testing.T) {
	n := YearsTwo(5)
	if n != 35 {
		t.Error("expected", 35, "got", 5)
	}
}

func ExampleYears() {
	fmt.Println(Years(5))
	// Output:
	// 35
}
func ExampleYearsTwo() {
	fmt.Println(YearsTwo(5))
	// Output:
	// 35
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(5)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(5)
	}
}
