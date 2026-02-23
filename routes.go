package main

import (
	"net/http"
)

func (a *app) routes() http.Handler {
	mux := http.NewServeMux()

	// fileServer := http.FileServer(http.Dir("/static/"))
	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.Handle("POST /download/poison_ivy_slayers_handbook_second_edition.pdf", downloadPdf(fileServer))
	mux.Handle("/", fileServer)

	return mux
}

func downloadPdf(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", "attachment")
		next.ServeHTTP(w, r)
	})
}
