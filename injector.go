//go:build wireinject
// +build wireinject

package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/middleware"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	controller.NewCategoryController,
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}
