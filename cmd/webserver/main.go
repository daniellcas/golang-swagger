package main

import (
	"net/http"

	"github.com/daniellcas/golang-swagger/internal/infra/webserver/handlers"
)

func main() {
	// Creating WebServer
	ws := http.NewServeMux()

	// Creating routes
	ws.HandleFunc("/finances", handlers.GetFinances)

	// Starting WebServer
	http.ListenAndServe(":3000", ws)
}
