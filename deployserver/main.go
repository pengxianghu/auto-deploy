package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func reLaunch()  {
	cmd := exec.Command("sh","../deploy.sh")
	err := cmd.Start()
	if nil != err {
		fmt.Println(err)
	}
	err = cmd.Wait()
	if nil != err {
		fmt.Println(err)
	}
}


func action(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>deploy server</h1>")
	reLaunch()
}

func main() {
	http.HandleFunc("/", action)
	http.ListenAndServe(":5001", nil)
}
