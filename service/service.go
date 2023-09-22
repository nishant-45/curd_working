package service

import (
	"context"
	"fmt"
	"sync"
	"taskOne/interfaces"
	"taskOne/models"
)

var service *Service
var serviceOnce sync.Once

type Service struct {
	repo interfaces.CategoryRepository
}

func NewService(repository interfaces.CategoryRepository) (service *Service) {
	serviceOnce.Do(func() {
		service = &Service{
			repo: repository,
		}
	})

	return service
}

func (sv *Service) CreateEntity(ctx context.Context, entity *models.Entity) (entity2 *models.Entity, err error) {
	//sv.repo.Create(entity)
	err = sv.repo.Create(ctx, entity)
	if err != nil {
		fmt.Printf("create error %v", err.Error())
		return
	}
	entity2 = entity
	return
}

func (sv *Service) ReadEntity(ctx context.Context) (entities []models.Entity, err error) {
	entities, err = sv.repo.Read(ctx)

	if err != nil {
		fmt.Printf("create error %v", err.Error())
		return
	}

	return entities, nil
}

func (sv *Service) DeleteEntity(ctx context.Context, UserId string) (err error) {
	err = sv.repo.Delete(ctx, UserId)
	fmt.Println(UserId)
	if err != nil {
		fmt.Printf("create error %v", err.Error())
		return
	}
	return
}

func (sv *Service) UpdateEntity(ctx context.Context, UserId string, entity *models.Entity) (err error) {
	err = sv.repo.Update(ctx, UserId, entity)

	if err != nil {
		fmt.Printf("create error %v", err.Error())
		return err
	}
	return nil
}
