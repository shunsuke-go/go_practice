package inputPort

import (
	"context"
	"go_practice/domain/models"
)

// usecase interactor (実質サービス層)のインターフェース

type IInitailUseCase interface {
	FindOneUser(ctx context.Context) (*models.User, error)
}