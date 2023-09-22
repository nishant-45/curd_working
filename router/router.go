package router

import (
	"taskOne/controller"
	"taskOne/mongodatabase"

	"github.com/gin-gonic/gin"
)

func Routes() {

	database := mongodatabase.Init().Database("test_db")
	ctrl := controller.Wire(database)

	r := gin.Default()
	r.POST("/entities", ctrl.CreateEntity)
	r.GET("/entities/", ctrl.ReadEntity)
	r.DELETE("/entities/:id", ctrl.DeleteEntity)
	r.PUT("/entities/:id", ctrl.UpdateEntity)
	r.Run(":8080")
}
