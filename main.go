package main

import (
	"dewaning/try-go-restful-api/app"
	"dewaning/try-go-restful-api/controller"
	"dewaning/try-go-restful-api/helper"
	"dewaning/try-go-restful-api/middleware"
	"dewaning/try-go-restful-api/repository"
	"dewaning/try-go-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
