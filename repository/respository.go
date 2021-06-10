package repsitory

import (
	"context"

	"github.com/Salauddin958/go-mux-product-apis/models"
)

// Product Repository explain...
type ProductRepo interface {
	Fetch(ctx context.Context, num int64) ([]*models.Product, error)
	GetByID(ctx context.Context, id int64) (*models.Product, error)
	Create(ctx context.Context, b *models.Product) (int64, error)
	Update(ctx context.Context, b *models.Product) (*models.Product, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
