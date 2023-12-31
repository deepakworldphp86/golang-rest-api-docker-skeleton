package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"
	"fmt"

	"github.com/deepakworldphp86/golang-api/graph/model"
)

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, categoryName string, description string) (*model.Categories, error) {
	panic(fmt.Errorf("not implemented: CreateCategory - createCategory"))
}

// SetMessage is the resolver for the setMessage field.
func (r *mutationResolver) SetMessage(ctx context.Context, message string) (string, error) {
		// Implement your mutation logic here
		return message, nil
}

// GetCategory is the resolver for the getCategory field.
func (r *queryResolver) GetCategory(ctx context.Context, id string) (*model.Categories, error) {
	panic(fmt.Errorf("not implemented: GetCategory - getCategory"))
}

// ListCategories is the resolver for the listCategories field.
func (r *queryResolver) ListCategories(ctx context.Context) ([]*model.Categories, error) {
	panic(fmt.Errorf("not implemented: ListCategories - listCategories"))
}

// Hello is the resolver for the hello field.
func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "Hello, GraphQL!", nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
