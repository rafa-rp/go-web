package routes

import (
	"go-web/controller"
	"net/http"
)

func LoadAllRoutes() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/insert", controller.SaveNewProduct)
	http.HandleFunc("/delete", controller.DeleteTableProduct)
}
