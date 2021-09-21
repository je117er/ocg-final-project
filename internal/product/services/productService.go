package services

import (
	"context"
	"github.com/je117er/ocg-final-project/internal/models"
	"github.com/je117er/ocg-final-project/internal/product"
)

type ProductService struct {
	productRepo product.Repository
}

func NewProductService(productRepo product.Repository) product.Service {
	return &ProductService{
		productRepo: productRepo,
		//ctx: ctx,
	}
}

func (ps *ProductService) ByID(ctx context.Context, id string) (*models.Product, error) {
	res, err := ps.productRepo.ByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ps *ProductService) All(ctx context.Context) ([]models.Product, error) {
	res, err := ps.productRepo.All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}