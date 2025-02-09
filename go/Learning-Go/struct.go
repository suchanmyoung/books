package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func struc() {
	var fred person
	bob := person{}
	julia := person{
		"Julia",
		40,
		"Cat",
	}

	fmt.Println(fred, bob, julia)
}

func main() {
	struc()
}
