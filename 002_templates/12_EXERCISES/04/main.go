package main

import (
	"log"
	"os"
	"text/template"
)

type menu struct {
	Categories []category
}

type category struct {
	Meal  string
	Foods []food
}

type food struct {
	Name  string
	Price float64
}

type restaurant struct {
	Name    string
	ResMenu menu
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r := []restaurant{
		restaurant{
			Name: "Old Pals Chicken",
			ResMenu: menu{
				Categories: []category{
					category{
						Meal: "Breakfast",
						Foods: []food{
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
					},
					category{
						Meal: "Lunch",
						Foods: []food{
							food{
								Name:  "Soup",
								Price: 3.9,
							},
							food{
								Name:  "Chicken Sandwich",
								Price: 4.5,
							},
							food{
								Name:  "Pizza slice",
								Price: 2.9,
							},
						},
					},
					category{
						Meal: "Dinner",
						Foods: []food{
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
					},
				},
			},
		},

		restaurant{
			Name: "Hasta la Vista Tacos",
			ResMenu: menu{
				Categories: []category{
					category{
						Meal: "Breakfast",
						Foods: []food{
							food{
								Name:  "Egg & Bacon Taco",
								Price: 4.9,
							},
							food{
								Name:  "Tomato Beans Taco",
								Price: 3.5,
							},
							food{
								Name:  "Taco Buffet",
								Price: 9.9,
							},
						},
					},
					category{
						Meal: "Lunch",
						Foods: []food{
							food{
								Name:  "Taco Soup",
								Price: 3.9,
							},
							food{
								Name:  "Chicken Tacos",
								Price: 4.5,
							},
							food{
								Name:  "Taco Pizza Slice",
								Price: 2.9,
							},
						},
					},
					category{
						Meal: "Dinner",
						Foods: []food{
							food{
								Name:  "Mega Uber Tacos",
								Price: 14.9,
							},
							food{
								Name:  "Fish Tacos",
								Price: 11.9,
							},
							food{
								Name:  "Taco Burger & Fries",
								Price: 14.9,
							},
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
