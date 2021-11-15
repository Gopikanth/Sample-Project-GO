package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	html := `<strong> Hello Kloudian </strong>`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, html)
}
func main() {
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":8080", nil)
}
