package handler

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter, r * http.Request, p httprouter.Params) {
	io.WriteString(w, " Create User String")

}