package interactor

import (
	"context"
	"go_practice/domain/entity"
	"go_practice/domain/repository"
	"go_practice/usecase/inputPort"
)

// usecase interactor (実質サービス層)の実装


type initialInteractor struct {
	repository repository.IInitialRepository
}

func NewInitialInteractor(repo repository.IInitialRepository) inputPort.IInitailUseCase {
	return &initialInteractor{
		repository: repo,
	}
}

func (ii *initialInteractor) FindOneUser (ctx context.Context) (*entity.User, error) {
	return ii.repository.FindOne(ctx)
}