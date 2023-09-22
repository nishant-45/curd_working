// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package controller

import (
	"go.mongodb.org/mongo-driver/mongo"
	"taskOne/repository"
	"taskOne/service"
)

// Injectors from wire.go:

func Wire(database *mongo.Database) *CategoryController {
	repositoryRepository := repository.NewCategoryRepo(database)
	serviceService := service.NewService(repositoryRepository)
	categoryController := NewController(serviceService)
	return categoryController
}
