package gateway

import (
	"context"
	"database/sql"
	"go_practice/domain/models"
	"go_practice/domain/repository"
)

type initialRepository struct {
	DB *sql.DB
}

func NewInitialRepository(db *sql.DB) repository.IInitialRepository {
	return &initialRepository{
		DB: db,
	}
}

func (ir *initialRepository) FindOne (ctx context.Context) (*models.User, error)  {
	return models.Users().One(ctx, ir.DB)
}