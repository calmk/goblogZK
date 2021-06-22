package main

import (
	"database/sql"
	"goblogCalmk/bootstrap"
	"goblogCalmk/pkg/database"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()
var db *sql.DB

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	database.Initialize()
	db = database.DB

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	router.Use(forceHTMLMiddleware)

	http.ListenAndServe(":8000", removeTrailingSlash(router))
}
