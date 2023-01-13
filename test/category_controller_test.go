package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"github.com/arifrachman98/go-restful-api/app"
	"github.com/arifrachman98/go-restful-api/controller"
	"github.com/arifrachman98/go-restful-api/helper"
	"github.com/arifrachman98/go-restful-api/middleware"
	"github.com/arifrachman98/go-restful-api/model/domain"
	"github.com/arifrachman98/go-restful-api/repository"
	"github.com/arifrachman98/go-restful-api/service"
	"github.com/go-playground/validator/v10"
)

var target = "http://localhost:3000/api/categories"

func setupHeader(r *http.Request) {
	r.Header.Add("Content-Type", "application-json")
	r.Header.Add("X-API-Key", "RAHASIA")
}

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_restful_api_test")
	helper.PanicHelper(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()

	cRepos := repository.NewCategoryRepository()
	cService := service.NewCategoryService(cRepos, db, validate)
	cController := controller.NewCategoryController(cService)

	router := app.NewRouter(cController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	rBody := strings.NewReader(`{"name" : "Gadget"}`)
	r := httptest.NewRequest(http.MethodPost, target, rBody)
	setupHeader(r)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, r)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	rBody := strings.NewReader(`{"name" : ""}`)
	r := httptest.NewRequest(http.MethodPost, target, rBody)
	setupHeader(r)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, r)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	cRepos := repository.NewCategoryRepository()
	c := cRepos.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	rBody := strings.NewReader(`{"name" : "Gadget"}`)
	r := httptest.NewRequest(http.MethodPut, target+"/"+strconv.Itoa(c.Id), rBody)
	setupHeader(r)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, r)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, c.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	cRepos := repository.NewCategoryRepository()
	c := cRepos.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	rBody := strings.NewReader(`{"name" : ""}`)
	r := httptest.NewRequest(http.MethodPut, target+"/"+strconv.Itoa(c.Id), rBody)
	setupHeader(r)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, r)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])
}

func TestGetCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	cRepos := repository.NewCategoryRepository()
	c := cRepos.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	r := httptest.NewRequest(http.MethodGet, target+"/"+strconv.Itoa(c.Id), nil)
	r.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, r)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, c.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, c.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	req := httptest.NewRequest(http.MethodGet, target+"/404", nil)
	req.Header.Add("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	res := rec.Result()
	assert.Equal(t, 404, res.StatusCode)

	body, _ := io.ReadAll(res.Body)

	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, 404, int(resBody["code"].(float64)))
	assert.Equal(t, "Not Found", resBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	cRepos := repository.NewCategoryRepository()
	c := cRepos.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	r := httptest.NewRequest(http.MethodDelete, target+"/"+strconv.Itoa(c.Id), nil)
	setupHeader(r)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, r)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	r := httptest.NewRequest(http.MethodDelete, target+"/404", nil)
	setupHeader(r)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, r)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
}

func TestListCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	cRepos := repository.NewCategoryRepository()
	c1 := cRepos.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	c2 := cRepos.Save(context.Background(), tx, domain.Category{
		Name: "Gojet",
	})
	tx.Commit()

	router := setupRouter(db)

	r := httptest.NewRequest(http.MethodGet, target, nil)
	r.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, r)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	var categories = responseBody["data"].([]interface{})
	categRes1 := categories[0].(map[string]interface{})
	categRes2 := categories[1].(map[string]interface{})

	assert.Equal(t, c1.Id, int(categRes1["id"].(float64)))
	assert.Equal(t, c2.Id, int(categRes2["id"].(float64)))

	assert.Equal(t, c1.Name, categRes1["name"])
	assert.Equal(t, c2.Name, categRes2["name"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	r := httptest.NewRequest(http.MethodGet, target, nil)
	r.Header.Add("X-API-Key", "Dede")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, r)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "Unauthorized", responseBody["status"])
}
