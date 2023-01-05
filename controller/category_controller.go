package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Param)
	Update(w http.ResponseWriter, r *http.Request, p httprouter.Param)
	Delete(w http.ResponseWriter, r *http.Request, p httprouter.Param)
	FindById(w http.ResponseWriter, r *http.Request, p httprouter.Param)
	FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Param)
}
