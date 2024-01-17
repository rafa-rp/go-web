package routes

import (
	"go-web/controller"
	"net/http"
)

func LoadAllRoutes() {
	http.HandleFunc("/", controller.Index)
}
