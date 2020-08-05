package main

import (
	"log"
	"os"
	"text/template"
)

type menu struct {
	Breakfast, Lunch, Dinner []food
}

type food struct {
	Name  string
	Price float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r := menu{
		Breakfast: []food{
			food{
				Name:  "Egg & Bacon",
				Price: 4.9,
			},
			food{
				Name:  "Tomato Beans",
				Price: 3.5,
			},
			food{
				Name:  "Buffet",
				Price: 9.9,
			},
		},
		Lunch: []food{
			food{
				Name:  "Soup",
				Price: 3.9,
			},
			food{
				Name:  "Sandwich",
				Price: 4.5,
			},
			food{
				Name:  "Pizza slice",
				Price: 2.9,
			},
		},
		Dinner: []food{
			food{
				Name:  "Taco",
				Price: 14.9,
			},
			food{
				Name:  "Salmon",
				Price: 11.9,
			},
			food{
				Name:  "Burger & Fries",
				Price: 14.9,
			},
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
