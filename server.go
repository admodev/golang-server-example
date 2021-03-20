package main

import (
	"net/http"
)

func main() {
	// Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contacts", contactHandler)

	// Start server
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is on."))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Contacts page"))
}
