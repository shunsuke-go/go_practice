package inputPort

import (
	"context"
	"go_practice/domain/entity"
)

// usecase interactor (実質サービス層)のインターフェース

type IInitailUseCase interface {
	FindOneUser(ctx context.Context) (*entity.User, error)
}