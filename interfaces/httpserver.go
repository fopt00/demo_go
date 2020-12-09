package interfaces

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HttpServerEntry() {
	dbPointer := &Database{
		"Allen": 19,
		"Luffy": 18,
		"Zoro":  20,
		"Sanji": 19,
	}
	log.Println("Server started at port:8001")
	err := http.ListenAndServe(":8001", dbPointer)
	log.Fatal(err)
}

type Database map[string]int

func (d *Database) ServeHTTP(w http.ResponseWriter, r *http.Request) {


	switch r.URL.Path {
	case "/":
		io.WriteString(w, "Hello world\n")
	case "/list":
		d.list(w, r)
	case "/age":
		d.age(w, r)
	default:
		io.WriteString(w, "Unknown path\n")
	}
}

func (d *Database) list(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html>\n<body>\n<ul>\n")
	for name, age := range *d {
		fmt.Fprintf(w, "<li>%s: %d</li>", name, age)
	}
	fmt.Fprintf(w, "</ul>\n</body>\n</html>\n")
}

func (d *Database) age(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age, exists := (*d)[name]
	if (exists) {
		fmt.Fprintf(w, "%s is %d years old\n", name, age)
	} else {
		fmt.Fprintf(w, "Cannot find %s's age\n", name)
	}
}
