package main

import (
	"GoBasic/go-todo-api/controller"
	"GoBasic/go-todo-api/model"
	"net/http"

	_ "github.com/go-sql-driver/sql" // mysql driver
)

func main() {

	mux := controller.Register()
	db := model.Connect()
	defer db.Close() // run when finished

	http.ListenAndServe(":3000", mux)

}
