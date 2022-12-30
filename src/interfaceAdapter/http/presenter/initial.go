package presenter

import (
	"go_practice/domain/entity"
	"go_practice/usecase/outputPort"
)

type initialPresenter struct {
	User *entity.User
}

func NewInitialPresenter(user *entity.User) outputport.IInitailPresenter {
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
