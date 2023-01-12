package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/arifrachman98/go-restful-api/helper"
	"github.com/arifrachman98/go-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(c context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "insert into category(name) values (?)"
	res, err := tx.ExecContext(c, sql, category.Name)
	helper.PanicHelper(err)

	id, err := res.LastInsertId()
	helper.PanicHelper(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(c context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(c, sql, category.Name, category.Id)
	helper.PanicHelper(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(c context.Context, tx *sql.Tx, category domain.Category) {
	sql := "delete from category where id = ?"
	_, err := tx.ExecContext(c, sql, category.Id)
	helper.PanicHelper(err)
}

func (repository *CategoryRepositoryImpl) FindById(c context.Context, tx *sql.Tx, categoryID int) (domain.Category, error) {
	sql := "select id, name from category where id = ?"
	rows, err := tx.QueryContext(c, sql, categoryID)
	helper.PanicHelper(err)
	defer rows.Close()

	categ := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&categ.Id, &categ.Name)
		helper.PanicHelper(err)
		return categ, nil
	} else {
		return categ, errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(c context.Context, tx *sql.Tx) []domain.Category {
	sql := "select id, name from category"
	rows, err := tx.QueryContext(c, sql)
	helper.PanicHelper(err)
	defer rows.Close()

	var categories []domain.Category

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicHelper(err)
		categories = append(categories, category)
	}

	return categories
}
