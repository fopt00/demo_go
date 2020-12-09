package complex_types

import (
	"html/template"
	"os"
)

func TemplateEntry() {
	//template1()
	template2()
}

func template1()  {
	m := Movie{
		Title: "Casablanca",
		Year:  1942,
		Color: false,
		Actors: []string{
			"Humphrey Bogart",
			"Ingrid Bergman",
		},
		duration: 1000,
	}

	tmpl, _ := template.New("text").Parse("{{.Title}} release in {{.Year}}.\n")
	tmpl.Execute(os.Stdout, m)
}

func template2()  {
	movies := []Movie{
		{"Casablanca", 1942, false, []string{"Humphrey Bogart", "Ingrid Bergman"}, 1000},
		{"Cool Hand Luke", 1967, true, []string{"Paul Newman"}, 1258},
		{"Bullitt", 1968, true, []string{"Steve McQueen", "Jacqueline Bisset"}, 2389},
	}

	tmpl, _ := template.New("text").
		Parse("{{range .}}{{.Title}} release in {{.Year}}.\n{{end}}")

	tmpl.Execute(os.Stdout, movies)
}
