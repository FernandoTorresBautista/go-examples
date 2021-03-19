package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	// functions handlers
	mux.HandleFunc("/ping", ping())

	return mux
}
