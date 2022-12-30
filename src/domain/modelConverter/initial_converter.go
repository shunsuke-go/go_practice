package modelconverter

import (
	"go_practice/domain/entity"
	"go_practice/domain/models"
)

func InitialConverter(user *models.User) *entity.User {
	return &entity.User{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}
}