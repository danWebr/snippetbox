package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Create a file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Register file server as handler
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Register the handler functions
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return app.logRequest(secureHeaders(mux))
}
