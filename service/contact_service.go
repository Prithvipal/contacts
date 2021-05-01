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

func GetContact(ctx context.Context, searchParam string) (contacts []dto.Contact, err error) {
	conts, err := dal.GetContact(ctx, searchParam)
	if err != nil {
		return
	}
	for _, cont := range conts {
		contact := convertToContactDto(cont)
		contacts = append(contacts, contact)
	}
	return contacts, nil

}

func PutContact(ctx context.Context, id string, cont dto.Contact) error {
	contact := convertToContactEntity(cont)
	return dal.PutContact(ctx, id, contact)
}

func DeleteContact(ctx context.Context, id string) error {
	return dal.DeleteContact(ctx, id)
}

func GetByIdContantsHandler(ctx context.Context, id string) (contact dto.Contact, err error) {
	cont, err := dal.GetByIdContantsHandler(ctx, id)
	if err != nil {
		return
	}
	return convertToContactDto(cont), nil
}

func convertToContactEntity(cont dto.Contact) (contact entity.Contact) {
	contact.Address = cont.Address
	contact.Name = cont.Name
	contact.Owners = cont.Owners
	contact.PhoneNumber = cont.PhoneNumber
	return
}

func convertToContactDto(cont entity.Contact) (contact dto.Contact) {
	contact.Id = cont.Id.Hex()
	contact.Address = cont.Address
	contact.Name = cont.Name
	contact.Owners = cont.Owners
	contact.PhoneNumber = cont.PhoneNumber
	return
}
