package mystr

import (
	"fmt"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {

	// const s = "Right now, the truth is declaring that all such situations like death, crime, penury, adversities, unemployment, disasters, maladies, depression, politics, and ignorance should not be classed as problems because problems do not exist at all, and they never did, only a problem. Nonetheless, you can find that you have already opted to be totally surprised, how it can be this way."
	s := "mahadev godbole"
	xs := strings.Split(s, " ")
	r := Cat(xs)

	if r != "mahadev godbole" {
		t.Error("got", r, "expected", s)
	}
}

func TestJoin(t *testing.T) {

	const s = "Right now, the truth is declaring that all such situations like death, crime, penury, adversities, unemployment, disasters, maladies, depression, politics, and ignorance should not be classed as problems because problems do not exist at all, and they never did, only a problem. Nonetheless, you can find that you have already opted to be totally surprised, how it can be this way."
	xs := strings.Split(s, " ")
	r := Join(xs)
	if r != s {
		t.Error("got", r, "expected", s)
	}
}

func ExampleCat() {
	// const s = `Right now, the truth is declaring that all such situations like death, crime, penury, adversities, unemployment, disasters, maladies, depression, politics, and ignorance should not be classed as problems because problems do not exist at all, and they never did, only a problem. Nonetheless, you can find that you have already opted to be totally surprised, how it can be this way.`
	const s = "Right now, the truth is declaring that all such situations like death, crime, penury, adversities, unemployment, disasters, maladies, depression, politics, and ignorance should not be classed as problems because problems do not exist at all, and they never did, only a problem. Nonetheless, you can find that you have already opted to be totally surprised, how it can be this way."
	xs := strings.Split(s, " ")
	fmt.Println(Cat(xs))
	// Output:
	// Right now, the truth is declaring that all such situations like death, crime, penury, adversities, unemployment, disasters, maladies, depression, politics, and ignorance should not be classed as problems because problems do not exist at all, and they never did, only a problem. Nonetheless, you can find that you have already opted to be totally surprised, how it can be this way.

}

func ExampleJoin() {
	// const s = `Right now, the truth is declaring that all such situations like death, crime, penury, adversities, unemployment, disasters, maladies, depression, politics, and ignorance should not be classed as problems because problems do not exist at all, and they never did, only a problem. Nonetheless, you can find that you have already opted to be totally surprised, how it can be this way.`
	const s = "Right now, the truth is declaring that all such situations like death, crime, penury, adversities, unemployment, disasters, maladies, depression, politics, and ignorance should not be classed as problems because problems do not exist at all, and they never did, only a problem. Nonetheless, you can find that you have already opted to be totally surprised, how it can be this way."
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs))
	// Output:
	// Right now, the truth is declaring that all such situations like death, crime, penury, adversities, unemployment, disasters, maladies, depression, politics, and ignorance should not be classed as problems because problems do not exist at all, and they never did, only a problem. Nonetheless, you can find that you have already opted to be totally surprised, how it can be this way.

}

func BenchmarkCat(b *testing.B) {
	const s = "Right now, the truth is declaring that all such situations like death, crime, penury, adversities, unemployment, disasters, maladies, depression, politics, and ignorance should not be classed as problems because problems do not exist at all, and they never did, only a problem. Nonetheless, you can find that you have already opted to be totally surprised, how it can be this way."
	xs := strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}

}

func BenchmarkJoin(b *testing.B) {
	const s = `Right now, the truth is declaring that all such situations like death, crime, penury, adversities, unemployment, disasters, maladies, depression, politics, and ignorance should not be classed as problems because problems do not exist at all, and they never did, only a problem. Nonetheless, you can find that you have already opted to be totally surprised, how it can be this way.`
	xs := strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(xs)
	}

}
