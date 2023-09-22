package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"sync"
	"taskOne/models"
)

var repository *Repository
var repoOnce sync.Once

type Repository struct {
	database *mongo.Database
}

func NewCategoryRepo(db *mongo.Database) *Repository {

	repoOnce.Do(func() {
		repository = &Repository{
			database: db,
		}
	})
	return repository
}

func (r *Repository) Create(ctx context.Context, entity *models.Entity) (err error) {

	_, err = r.database.Collection("entities").InsertOne(ctx, entity)
	if err != nil {
		fmt.Printf("create error %v", err.Error())
		return
	}
	return
}

func (r *Repository) Read(ctx context.Context) (users []models.Entity, err error) {
	var entities []models.Entity
	cursor, err := r.database.Collection("entities").Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	//defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var entity models.Entity
		if err := cursor.Decode(&entity); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

func (r *Repository) Delete(ctx context.Context, UserId string) (err error) {
	fmt.Println(UserId)
	data, err := strconv.ParseInt(UserId, 10, 32)
	if err != nil {
		fmt.Println("err", err.Error())
		return err
	}

	filter := bson.M{"id": data}

	_, err = r.database.Collection("entities").DeleteOne(context.TODO(), filter)

	if err != nil {
		fmt.Println("err", err.Error())
		return err
	}
	//fmt.Printf("%v", res.DeletedCount)
	return nil
}

func (r *Repository) Update(ctx context.Context, UserId string, entity *models.Entity) (err error) {
	data, err := strconv.ParseInt(UserId, 10, 32)

	if err != nil {
		fmt.Println("err", err.Error())
		return err
	}

	filter := bson.M{"id": data}
	update := bson.M{"$set": entity}

	_, err = r.database.Collection("entities").UpdateOne(context.TODO(), filter, update)

	if err != nil {
		fmt.Println("err", err.Error())
		return err
	}

	return nil
}
