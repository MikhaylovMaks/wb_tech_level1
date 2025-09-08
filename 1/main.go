package main

import "fmt"

type Human struct {
	Name         string
	Surname      string
	Age          int
	PlaceOfBirth string
}

type Action struct {
	Human
	Active string
}

func (h Human) DateBirthday() {
	fmt.Printf("%s %s родился в городе %s, ему %d лет\n", h.Name, h.Surname, h.PlaceOfBirth, h.Age)
}

func (a Action) DoSomething() {
	fmt.Printf("Что происходит с %s: %s\n", a.Name, a.Active)
}

func main() {
	human := Human{
		Name:         "Maksim",
		Surname:      "Mikhaylov",
		Age:          20,
		PlaceOfBirth: "Kazan",
	}
	a := Action{
		Human:  human,
		Active: "взрослеет",
	}
	human.DateBirthday()
	a.DateBirthday()
	a.DoSomething()
	// нельзя human.DoSomething()
}
