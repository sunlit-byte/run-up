package main

import (
	"RunUp/handler"
	"net/http"

	"github.com/julienschmidt/httprouter"

)

func main() {

	r := RegisterHandler()

	http.ListenAndServe(":8080",r)

}


func RegisterHandler() *httprouter.Router{
	router := httprouter.New()

	router.POST("/user", handler.CreateUser)

	return router
}

