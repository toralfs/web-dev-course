package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

type region struct {
	Region string
	Hotels []hotel
}

// Alternatively: Create a regions type, but this feels a bit unnecessary to me:
// type regions []region

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r := []region{
		region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
				},
				hotel{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
				},
			},
		},
		region{
			Region: "Northern",
			Hotels: []hotel{
				hotel{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
					Region:  "southern",
				},
				hotel{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
					Region:  "southern",
				},
			},
		},
		region{
			Region: "Central",
			Hotels: []hotel{
				hotel{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
					Region:  "southern",
				},
				hotel{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
					Region:  "southern",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
