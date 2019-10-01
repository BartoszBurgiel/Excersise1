package main

import (
	"excersises/dailywalk/person"
	"fmt"
)

func main() {
	fmt.Println("hello")

	p := person.NewPerson("Bob")

	p.TideShoes()
}
