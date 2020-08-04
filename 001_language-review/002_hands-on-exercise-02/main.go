package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) pSpeak() {
	fmt.Println("This is person speaking")
}

func (sa secretAgent) saSpeak() {
	fmt.Println("This is secret agent speaking")
}

func main() {
	p := person{
		fname: "Dolly",
		lname: "Duck",
		age:   30,
	}

	sa := secretAgent{
		person: person{
			fname: "Fantonald",
			lname: "Duck",
			age:   32,
		},
		ltk: false,
	}

	fmt.Printf("Hi, I am %s %s, I am %d years old\n", p.fname, p.lname, p.age)
	p.pSpeak()

	fmt.Printf("SHHH, be quiet, my name is %s, I gonna get these bad guys.\n", sa.fname)
	sa.saSpeak()
	sa.pSpeak()
}
