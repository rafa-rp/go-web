package main

import (
	"go-web/routes"
	"net/http"
)

func main() {
	routes.LoadAllRoutes()
	http.ListenAndServe(":8080", nil)
}
