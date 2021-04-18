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

func convertToContactEntity(cont dto.Contact) (contact entity.Contact) {
	contact.Address = cont.Address
	contact.Name = cont.Name
	contact.Owners = cont.Owners
	contact.PhoneNumber = cont.PhoneNumber
	return
}
