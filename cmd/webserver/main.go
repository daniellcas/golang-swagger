package main

import (
	"net/http"

	_ "github.com/daniellcas/golang-swagger/docs"
	"github.com/daniellcas/golang-swagger/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

func authMiddware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if len(auth) == 0 {
			http.Error(w, "please set Header Authorization", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

//	@title			Finances example api
//	@version		v0.0.1
//	@description	This is an example of an api using swagger

//	@contact.name	Daniel lucas
//	@contact.email	danielsantos120615@gmail.com

//	@BasePath	/

//	@securitydefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
func main() {
	// Creating WebServer
	ws := http.NewServeMux()

	// Creating routes
	ws.HandleFunc("GET /finances", authMiddware(handlers.GetFinances))

	// Docs
	ws.HandleFunc("GET /docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:3000/docs/doc.json")))

	ws.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs", http.StatusPermanentRedirect)
	})

	// Starting WebServer
	http.ListenAndServe(":3000", ws)
}
