package dto

type Contact struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Address     string   `json:"address"`
	PhoneNumber []int    `json:"phone_number"`
	Owners      []string `json:"owners"`
}
