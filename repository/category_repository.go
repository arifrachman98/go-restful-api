package repository

import (
	"context"
	"database/sql"

	"github.com/arifrachman98/go-restful-api/model/domain"
)

type CategoryRepository interface {
	Save(c context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(c context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(c context.Context, tx *sql.Tx, category domain.Category)
	FindById(c context.Context, tx *sql.Tx, categoryID int) domain.Category
	FindAll(c context.Context, tx *sql.Tx) []domain.Category
}
