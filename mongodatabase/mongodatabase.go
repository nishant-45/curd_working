package mongodatabase

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Init() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://Nishant_45:Nishant4321@cluster0.bsdl3tr.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

