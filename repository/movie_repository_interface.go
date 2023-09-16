package repository

import (
	"api-mysql/models"
	"context"
)

type MovieRepositoryInterface interface {
	GetAll(ctx context.Context) ([]models.Movie, error)
	Insert(ctx context.Context, movie models.Movie) error
	Update(ctx context.Context, movie models.Movie, id string) error
	Delete(ctx context.Context, id string) error
}
