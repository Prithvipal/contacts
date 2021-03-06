package dal

import (
	"context"
	"fmt"

	"github.com/Prithvipal/phone-dir/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetContact(ctx context.Context) (contacts []entity.Contact, err error) {

	client, err := connect(ctx)
	if err != nil {
		return
	}
	defer client.Disconnect(ctx)

	contactDB := client.Database("contact")
	companyColl := contactDB.Collection("companies")

	cursor, err := companyColl.Find(ctx, bson.M{})
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &contacts); err != nil {
		return
	}

	return
}
func FindNameContains(ctx context.Context, searchParam string) (contacts []entity.Contact, err error) {
	client, err := connect(ctx)
	if err != nil {
		return
	}
	defer client.Disconnect(ctx)

	contactDB := client.Database("contact")
	companyColl := contactDB.Collection("companies")
	cursor, err := companyColl.Find(ctx, bson.M{
		"name": primitive.Regex{
			Pattern: searchParam,
			Options: "i",
		},
	})
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &contacts); err != nil {
		return
	}

	return
}

func PutContact(ctx context.Context, id string, cont entity.Contact) error {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	client, err := connect(ctx)
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	contactDB := client.Database("contact")
	companyColl := contactDB.Collection("companies")

	update := bson.M{
		"$set": cont,
	}
	upsert := true
	opt := options.FindOneAndUpdateOptions{Upsert: &upsert}
	result := companyColl.FindOneAndUpdate(ctx, bson.M{"_id": objID}, update, &opt)
	return result.Err()

}

func DeleteContact(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	client, err := connect(ctx)
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	contactDB := client.Database("contact")
	companyColl := contactDB.Collection("companies")
	result, err := companyColl.DeleteOne(ctx, bson.M{"_id": objID})
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
	return err

}
func GetByIdContantsHandler(ctx context.Context, id string) (contact entity.Contact, err error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	client, err := connect(ctx)
	if err != nil {
		return
	}
	defer client.Disconnect(ctx)

	contactDB := client.Database("contact")
	companyColl := contactDB.Collection("companies")
	err = companyColl.FindOne(ctx, bson.M{"_id": objID}).Decode(&contact)
	return

}
