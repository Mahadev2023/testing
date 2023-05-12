package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	c := Count("Ma dev god")
	if c != 3 {
		t.Error("expected", 3, "got", c)
	}
}
func TestUseCount(t *testing.T) {
	c := UseCount("Ma dev god")

	ans := map[string]int{"Ma": 1, "dev": 1, "god": 1}
	for k, v := range c {
		if ans[k] != v {
			t.Error("expected", ans, "got", c)
			break
		}
	}
}

func ExampleCount() {
	fmt.Println(Count("Ma dev god"))
	// Output
	//3
}

func ExampleUseCount() {
	fmt.Println(UseCount("Ma dev god"))
	//Output
	//map[Ma:1 dev:1 god:1]
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("Ma dev god")
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("Ma dev god")
	}
}
