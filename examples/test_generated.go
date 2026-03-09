package examples

import (
	"fmt"
)

func TestGenerated() {
	p := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			City:    "Moscow",
			Country: "Russia",
		},
	}
	m := PersonToMap(p)
	fmt.Printf("%#v\n", m)
}
