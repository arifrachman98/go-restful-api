package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/arifrachman98/go-restful-api/helper"
	"github.com/arifrachman98/go-restful-api/model/web"
	"github.com/arifrachman98/go-restful-api/service"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func (c *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	categoryCreateRequest := web.CategoryCreateRequest{}
	err := decoder.Decode(&categoryCreateRequest)
	helper.PanicHelper(err)

	categoryResponse := c.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicHelper(err)
}

func (c *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	catgoryUpdateRequest := web.CategoryUpdateRequest{}
	err := decoder.Decode(&catgoryUpdateRequest)
	helper.PanicHelper(err)

	categoryID := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryID)
	helper.PanicHelper(err)

	catgoryUpdateRequest.Id = id

	categoryResponse := c.CategoryService.Update(r.Context(), catgoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	endcode := json.NewEncoder(w)
	err = endcode.Encode(webResponse)
	helper.PanicHelper(err)
}

func (c *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryID := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryID)
	helper.PanicHelper(err)

	c.CategoryService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	w.Header().Add("Content-Type", "application/json")
	endcode := json.NewEncoder(w)
	err = endcode.Encode(webResponse)
	helper.PanicHelper(err)
}

func (c *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryID := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryID)
	helper.PanicHelper(err)

	categoryResponse := c.CategoryService.FindByID(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	endcode := json.NewEncoder(w)
	err = endcode.Encode(webResponse)
	helper.PanicHelper(err)
}

func (c *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryResponses := c.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	w.Header().Add("Content-Type", "application/json")
	endcode := json.NewEncoder(w)
	err := endcode.Encode(webResponse)
	helper.PanicHelper(err)
}
