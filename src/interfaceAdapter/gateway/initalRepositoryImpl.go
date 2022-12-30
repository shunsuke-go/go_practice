package gateway

import (
	"context"
	"database/sql"
	"fmt"
	"go_practice/domain/entity"
	modelconverter "go_practice/domain/modelConverter"
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

func (ir *initialRepository) FindOne (ctx context.Context) (*entity.User, error)  {
	userModel, err := models.Users().One(ctx, ir.DB)

	if err != nil {
		fmt.Print("db error", err)
		return nil, err
	}
	user := modelconverter.InitialConverter(userModel)
	return user, nil
}