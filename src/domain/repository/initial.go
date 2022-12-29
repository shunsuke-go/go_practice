package repository

import (
	"context"
	"go_practice/domain/models"
)

type IInitialRepository interface {
	FindOne(ctx context.Context) (*models.User, error)
}