//go:build wireinject
// +build wireinject

package main

import (
	"github.com/arifrachman98/go-restful-api/app"
	"github.com/arifrachman98/go-restful-api/controller"
	"github.com/arifrachman98/go-restful-api/middleware"
	"github.com/arifrachman98/go-restful-api/repository"
	"github.com/arifrachman98/go-restful-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitServer() *http.Server {
	wire.Build(app.NewDB, validator.New, categorySet, app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}
