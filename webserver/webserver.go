package main

import (
	"io"
	"net/http"
)

func action(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>qqqqqqq</h1>")
}


func main() {
	http.HandleFunc("/", action)
	http.ListenAndServe(":5000", nil)
}
