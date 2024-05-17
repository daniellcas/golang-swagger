package main

import (
	"net/http"

	_ "github.com/daniellcas/golang-swagger/docs"
	"github.com/daniellcas/golang-swagger/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Finances example api
// @version v0.0.1
// @description This is an example of an api using swagger

// @contact.name Daniel lucas
// @contact.email danielsantos120615@gmail.com

// @BasePath /
func main() {
	// Creating WebServer
	ws := http.NewServeMux()

	// Creating routes
	ws.HandleFunc("GET /finances", handlers.GetFinances)

	// Docs
	ws.HandleFunc("GET /docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:3000/docs/doc.json")))

	// Starting WebServer
	http.ListenAndServe(":3000", ws)
}
