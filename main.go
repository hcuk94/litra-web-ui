package main

import (
	"fmt"
	"net/http"

	"github.com/derickr/go-litra-driver"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, world!</h1>")
}

func main() {

	ld, err := litra.New()
	if err != nil {
		fmt.Println("I am sad :-(")
	}
	ld.TurnOn()
	ld.Close()
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting server...")
	http.ListenAndServe(":8080", nil)
}
