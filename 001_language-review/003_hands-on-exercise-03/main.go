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

func (p person) run() {
	fmt.Printf("%s %s is running, very fast!\n", p.fname, p.lname)
}

func (sa secretAgent) saSpeak() {
	fmt.Println("This is secret agent speaking")
}

func (sa secretAgent) run() {
	fmt.Printf("%s is running like the wind!\n", sa.fname)
}

type human interface {
	run()
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

	running(p)
	running(sa)
}

func running(h human) {
	h.run()
}
