package dal

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://prithvi:prithvi123@cluster0.rmlet.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = client.Connect(ctx)
	return client, err
}
