package service

import (
	"context"

	"github.com/arifrachman98/go-restful-api/model/web"
)

type CategoryService interface {
	Create(c context.Context, r web.CategoryCreateRequest) web.CategoryResponse
	Update(c context.Context, r web.CategoryUpdateRequest) web.CategoryResponse
	Delete(c context.Context, categoryID int)
	FindByID(c context.Context, categoryID int) web.CategoryResponse
	FindAll(c context.Context) []web.CategoryResponse
}
