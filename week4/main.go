package main

import (
	"fmt"
	"net/http"
)

type data struct {
	str string
}

func (d *data) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "handl")
}

func main() {
	d := &data{str: "name"}
	http.Handle("/", d)

	http.ListenAndServe(":8080", nil)
}
