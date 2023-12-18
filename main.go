package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name, Description string
	Price             float64
	Quantity          int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	Products := []Product{
		{"T-Shirt", "Nice Cloth", 90, 10},
		{"Shoes", "Confortable", 200, 6},
	}

	temp.ExecuteTemplate(w, "Index", Products)
}
