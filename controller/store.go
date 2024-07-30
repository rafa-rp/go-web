package controller

import (
	"go-web/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.FindAllProducts()

	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func SaveNewProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting the price", err)
		}

		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error converting the quantity", err)
		}

		models.SaveProduct(name, description, priceConverted, quantityConverted)
	}
	http.Redirect(w, r, "/", 301)
}
