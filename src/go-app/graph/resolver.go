package graph

import (
	"context"
	"github.com/deepakworldphp86/golang-api/graph/model"
)

type Resolver struct {
	repo model.CategoryRepository
}

func NewResolver(repo model.CategoryRepository) *Resolver {
	return &Resolver{repo}
}

func (r *Resolver) GetCategory(ctx context.Context, id uint) (*model.Categories, error) {
	return r.repo.GetCategoryByID(id)
}

func (r *Resolver) ListCategories(ctx context.Context) ([]*model.Categories, error) {
	return r.repo.ListCategories()
}

func (r *Resolver) CreateCategory(ctx context.Context, category_name string, description string) (*model.Categories, error) {
	return r.repo.CreateCategory(category_name, description)
}

