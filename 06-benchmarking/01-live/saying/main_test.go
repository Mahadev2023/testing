package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Mahadev")

	if s != "Welcome my dear Mahadev" {
		t.Error("got", s, "expected", "Welcome my dear Mahadev")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Mahadev"))
	// Output:
	// Welcome my dear Mahadev
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Mahadev")
	}
}
