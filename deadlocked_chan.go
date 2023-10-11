package main

import "fmt"

type person struct {
	name      string
	forkLeft  string
	forkRight string
}

type fork struct {
	status string
}

func newPerson(name, forkLeft, forkRight string) *person {

	p := person{name: name, forkLeft: forkLeft, forkRight: forkRight}

	return &p
}

func newFork(status string) *fork {
	f := fork{
		status: status,
	}
	return &f
}

func main() {

	fmt.Println(person{"Bob", "inuse", "inuse"})
	fmt.Println(person{"Per", "inuse", "available"})
	fmt.Println(person{"Poul", "inuse", "inuse"})

}
