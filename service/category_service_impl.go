package service

import (
	"context"
	"database/sql"

	"github.com/arifrachman98/go-restful-api/helper"
	"github.com/arifrachman98/go-restful-api/model/domain"
	"github.com/arifrachman98/go-restful-api/model/web"
	"github.com/arifrachman98/go-restful-api/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

func (service *CategoryServiceImpl) Create(c context.Context, r web.CategoryCreateRequest) web.CategoryResponse {
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
	tx, err := service.DB.Begin()
	helper.PanicHelper(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(c, tx, r.Id)
	helper.PanicHelper(err)

	category = service.CategoryRepository.Update(c, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(c context.Context, categoryID int) {
	tx, err := service.DB.Begin()
	helper.PanicHelper(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(c, tx, categoryID)
	helper.PanicHelper(err)

	service.CategoryRepository.Delete(c, tx, category)
}

func (service *CategoryServiceImpl) FindByID(c context.Context, categoryID int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicHelper(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(c, tx, categoryID)
	helper.PanicHelper(err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(c context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicHelper(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(c, tx)

	return helper.ToCategoryResponses(categories)
}
