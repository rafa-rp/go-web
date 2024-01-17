package controller

import (
	"go-web/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.FindAllProducts()

	temp.ExecuteTemplate(w, "Index", products)
}
