package repository

import (
	"context"
	"github.com/trisatya23/go-restful-api-master/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, category domain.Category) (domain.Category, error)
	Update(ctx context.Context, category domain.Category) (domain.Category, error)
	Delete(ctx context.Context, category domain.Category) error
	FindById(ctx context.Context, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context) ([]domain.Category, error)
}
