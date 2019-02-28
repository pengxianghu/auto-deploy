package main

import (
	"io"
	"net/http"
)

func action(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>1234567890</h1>")
}


func main() {
	http.HandleFunc("/", action)
	http.ListenAndServe(":5000", nil)
}
