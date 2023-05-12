package main

import (
	"fmt"
	"strings"

	"github.com/Mahadev2023/testing/07-Benchamarking-example1/mystr"
)

const s = `Right now, the truth is declaring that all such situations like death, crime, penury, adversities, unemployment, disasters, maladies, depression, politics, and ignorance should not be classed as problems because problems do not exist at all, and they never did, only a problem. Nonetheless, you can find that you have already opted to be totally surprised, how it can be this way.`

var xs []string

func main() {
	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}
	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n", mystr.Join(xs))
}
