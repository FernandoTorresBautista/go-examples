package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("inicializando el server")

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received")
		fmt.Println(r.Method)
		w.Write([]byte("Hello world"))
	})

	http.ListenAndServe("localhost:3000", mux)

}
