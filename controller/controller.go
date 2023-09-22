package controller

import (
	"fmt"
	"net/http"
	"sync"
	"taskOne/interfaces"
	"taskOne/models"

	"github.com/gin-gonic/gin"
)

var controllerOnce sync.Once
var controller *CategoryController

type CategoryController struct {
	service interfaces.CategoryService
}

func NewController(categoryService interfaces.CategoryService) *CategoryController {
	controllerOnce.Do(func() {
		controller = &CategoryController{
			service: categoryService,
		}
	})

	return controller
}

func (ctrl *CategoryController) CreateEntity(c *gin.Context) {
	var entity *models.Entity
	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%v", entity)

	entity, err := ctrl.service.CreateEntity(c, entity)
	if err != nil {
		fmt.Printf("create error %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, entity)
}

func (ctrl *CategoryController) ReadEntity(c *gin.Context) {
	entity, err := ctrl.service.ReadEntity(c)

	if err != nil {
		fmt.Printf("create error %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, entity)
}

func (ctrl *CategoryController) DeleteEntity(c *gin.Context) {
	userId := c.Param("id")

	fmt.Println(userId)

	err := ctrl.service.DeleteEntity(c, userId)

	if err != nil {
		fmt.Printf("create error %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (ctrl *CategoryController) UpdateEntity(c *gin.Context) {

	userId := c.Param("id")

	var entity *models.Entity
	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%v", entity)

	err := ctrl.service.UpdateEntity(c, userId, entity)

	if err != nil {
		fmt.Printf("create error %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Entity updated successfully"})
}
