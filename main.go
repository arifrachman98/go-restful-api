package main

import (
	"github.com/arifrachman98/go-restful-api/app"
	"github.com/arifrachman98/go-restful-api/controller"
	"github.com/arifrachman98/go-restful-api/repository"
	"github.com/arifrachman98/go-restful-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	cRepos := repository.NewCategoryRepository()
	cService := service.NewCategoryService(cRepos, db, validate)
	cController := controller.NewCategoryController(cService)

	router := httprouter.New()

	router.GET("/api/categories", cController.FindAll)
	router.GET("/api/categories/:categoryId", cController.FindById)
	router.POST("/api/categories", cController.Create)
	router.PUT("/api/categories/:categoryId", cController.Update)
	router.DELETE("/api/categories/:categoryId", cController.Delete)

}
