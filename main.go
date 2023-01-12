package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/arifrachman98/go-restful-api/app"
	"github.com/arifrachman98/go-restful-api/controller"
	"github.com/arifrachman98/go-restful-api/helper"
	"github.com/arifrachman98/go-restful-api/repository"
	"github.com/arifrachman98/go-restful-api/service"
	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	cRepos := repository.NewCategoryRepository()
	cService := service.NewCategoryService(cRepos, db, validate)
	cController := controller.NewCategoryController(cService)

	router := app.NewRouter(cController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicHelper(err)
}
