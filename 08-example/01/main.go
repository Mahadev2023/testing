package starting_code

import (
	"fmt"

	"github.com/Mahadev2023/testing/08-example/01/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}

/*
 go test
go test bench .
go test cover
go test -coverprofile c.out
go tool cover -html=c.out */
