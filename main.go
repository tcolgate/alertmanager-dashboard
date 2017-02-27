package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening.")
	http.Handle("/", http.FileServer(http.Dir("./build/bundled")))
	http.ListenAndServe(":3002", nil)
}
