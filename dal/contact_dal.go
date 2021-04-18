package dal

import (
	"context"

	"github.com/Prithvipal/phone-dir/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveContact(ctx context.Context, cont entity.Contact) error {
	client, err := connect(ctx)
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	contactDB := client.Database("contact")
	companyColl := contactDB.Collection("companies")

	cont.Id = primitive.NewObjectID()
	_, err = companyColl.InsertOne(ctx, cont)
	return err
}
