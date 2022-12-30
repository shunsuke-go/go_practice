package repository

import (
	"context"
	"go_practice/domain/entity"
)

type IInitialRepository interface {
	FindOne(ctx context.Context) (*entity.User, error)
}