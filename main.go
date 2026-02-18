package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
	fmt.Fprintln(w, "\nHelllow world")
}
func main() {

	http.HandleFunc("/", hello)

	fmt.Println("Server Started at port 8080")
	http.ListenAndServe(":8080", nil)
}
