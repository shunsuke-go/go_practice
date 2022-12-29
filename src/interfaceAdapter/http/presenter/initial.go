package presenter

import (
	"go_practice/domain/models"
	outputport "go_practice/usecase/outputPort"
)

type initialPresenter struct {
	User *models.User
}

func NewInitialPresenter(user *models.User) outputport.IInitailPresenter {
	return &initialPresenter{
		User: user,
	}
}

func (ip *initialPresenter) ToInitialPresenter() outputport.InitialPresenter  {
	return outputport.InitialPresenter{
		Id: ip.User.ID,
		Name: ip.User.Name,
		Email: ip.User.Email,
	}
}
