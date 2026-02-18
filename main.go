package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
	fmt.Fprintln(w, "\nHelllow world")
}

func sayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Heyy I am From Server")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/hey", sayHi)
	fmt.Println("Server Started at port 8080")
	http.ListenAndServe(":8080", mux)
}
