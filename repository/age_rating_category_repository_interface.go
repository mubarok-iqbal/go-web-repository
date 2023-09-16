package repository

import (
	"api-mysql/models"
	"context"
)

type AgeRatingCategoryRepositoryInterface interface {
	GetAll(ctx context.Context) ([]models.AgeRatingCategory, error)
	Insert(ctx context.Context, rating models.AgeRatingCategory) error
	Update(ctx context.Context, rating models.AgeRatingCategory, id string) error
	Delete(ctx context.Context, id string) error
}
