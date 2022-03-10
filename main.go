package main

import (
	"net/http"

	"github.com/DanielDutraMl/Go-Web/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
