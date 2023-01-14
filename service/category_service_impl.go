package service

import (
	"context"
	"database/sql"

	"github.com/arifrachman98/go-restful-api/exception"
	"github.com/arifrachman98/go-restful-api/helper"
	"github.com/arifrachman98/go-restful-api/model/domain"
	"github.com/arifrachman98/go-restful-api/model/web"
	"github.com/arifrachman98/go-restful-api/repository"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(cRepos repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		CategoryRepository: cRepos,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(c context.Context, r web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(r)
	helper.PanicHelper(err)

	tx, err := service.DB.Begin()
	helper.PanicHelper(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: r.Name,
	}

	category = service.CategoryRepository.Save(c, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(c context.Context, r web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(r)
	helper.PanicHelper(err)

	tx, err := service.DB.Begin()
	helper.PanicHelper(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(c, tx, r.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = r.Name

	category = service.CategoryRepository.Update(c, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(c context.Context, categoryID int) {
	tx, err := service.DB.Begin()
	helper.PanicHelper(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(c, tx, categoryID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(c, tx, category)
}

func (service *CategoryServiceImpl) FindByID(c context.Context, categoryID int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicHelper(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(c, tx, categoryID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(c context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicHelper(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(c, tx)

	return helper.ToCategoryResponses(categories)
}
