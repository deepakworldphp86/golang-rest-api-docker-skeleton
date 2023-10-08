package model

type CategoryRepository interface {
	GetCategoryByID(id uint) (*Categories, error)
	ListCategories() ([]*Categories, error)
	CreateCategory(categoryName, description string) (*Categories, error)
}
