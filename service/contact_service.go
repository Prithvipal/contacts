package service

import (
	"context"

	"github.com/Prithvipal/phone-dir/dal"
	"github.com/Prithvipal/phone-dir/dto"
	"github.com/Prithvipal/phone-dir/entity"
)

func SaveContact(ctx context.Context, cont dto.Contact) error {
	contact := convertToContactEntity(cont)
	return dal.SaveContact(ctx, contact)
}

func GetContact(ctx context.Context) ([]byte, error) {
	return dal.GetContact(ctx)
}

func PutContact(ctx context.Context, cont dto.Contact, id string) error {
	contact := convertToContactEntity(cont)
	return dal.PutContact(ctx, contact, id)
}

func DeleteContact(ctx context.Context, id string) error {
	return dal.DeleteContact(ctx, id)
}

func GetByIdContantsHandler(ctx context.Context, id string) ([]byte, error) {
	return dal.GetByIdContantsHandler(ctx, id)
}

func convertToContactEntity(cont dto.Contact) (contact entity.Contact) {
	contact.Address = cont.Address
	contact.Name = cont.Name
	contact.Owners = cont.Owners
	contact.PhoneNumber = cont.PhoneNumber
	return
}
