package main

import (
	"io"
	"net/http"
)

func action(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>blabla</h1>")
}


func main() {
	http.HandleFunc("/", action)
	http.ListenAndServe(":5000", nil)
}
