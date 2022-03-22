package main

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	//db := app.NewDB()
	//validate := validator.New()
	//categoryRepository := repository.NewCategoryRepository()
	//categoryService := service.NewCategoryService(categoryRepository, db, validate)
	//categoryController := controller.NewCategoryController(categoryService)
	//
	//router := app.NewRouter(categoryController)
	//authMiddleware := middleware.NewAuthMiddleware(router)

	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
