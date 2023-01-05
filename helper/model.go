package helper

import (
	"github.com/arifrachman98/go-restful-api/model/domain"
	"github.com/arifrachman98/go-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
