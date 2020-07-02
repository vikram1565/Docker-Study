package main

import (
	"io"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World, from docker container")
}
func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":80", nil)
}
