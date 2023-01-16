package app

import (
	"database/sql"
	"time"

	"github.com/arifrachman98/go-restful-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_database_migration")
	helper.PanicHelper(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

/*
	! Create Migrate
	* migrate create -ext sql -dir db/migrations create_table_category
*/

/*
	! Migrate up
	* migrate -database "mysql://root@tcp(localhost:3306)/go_database_migration" -path db/migrations up

	! Migrate down
	* * migrate -database "mysql://root@tcp(localhost:3306)/go_database_migration" -path db/migrations down
*/

/*
	! Check version of migrate state
	* migrate -database "mysql://root@tcp(localhost:3306)/go_database_migration" -path db/migrations version

	! Rollback version of migrate state
	* migrate -database "mysql://root@tcp(localhost:3306)/go_database_migration" -path db/migrations force 20230116175811
*/