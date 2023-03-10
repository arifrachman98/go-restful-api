package main

import (
	"fmt"
	"net/http"

	"github.com/arifrachman98/go-restful-api/helper"
	"github.com/arifrachman98/go-restful-api/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func NewServer(AuthMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: AuthMiddleware,
	}
}

func main() {
	server := InitializedServer()
	fmt.Printf("Listening on port :3000")
	err := server.ListenAndServe()
	helper.PanicHelper(err)
}
