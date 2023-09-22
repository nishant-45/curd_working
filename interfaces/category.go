package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
	"taskOne/models"
)

type CategoryController interface {
	CreateEntity(c *gin.Context)
}

type CategoryService interface {
	CreateEntity(ctx context.Context, entity *models.Entity) (entity2 *models.Entity, err error)
	ReadEntity(ctx context.Context) (entities []models.Entity, err error)
	DeleteEntity(ctx context.Context, UserId string) (err error)
	UpdateEntity(ctx context.Context, UserId string, entity *models.Entity) (err error)
}

type CategoryRepository interface {
	Create(ctx context.Context, entity *models.Entity) (err error)
	Read(ctx context.Context) (users []models.Entity, err error)
	Delete(ctx context.Context, UserId string) (err error)
	Update(ctx context.Context, UserId string, entity *models.Entity) (err error)
}
