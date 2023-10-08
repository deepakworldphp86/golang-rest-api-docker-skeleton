package repository

import (
	"github.com/deepakworldphp86/golang-api/entity"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	InsertCategory(category entity.Categories) entity.Categories
	UpdateCategory(category entity.Categories) entity.Categories
	GetCategory(categoryID uint64) entity.Categories
	DeleteCategory(category entity.Categories)
	AllCategory() []entity.Categories
}

type categoryConnection struct {
	connection *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryConnection{
		connection: db,
	}
}

func (db *categoryConnection) InsertCategory(category entity.Categories) entity.Categories {
	db.connection.Save(&category)
	return category
}

func (db *categoryConnection) UpdateCategory(category entity.Categories) entity.Categories {
	db.connection.Save(&category)
	return category
}

func (db *categoryConnection) GetCategory(categoryID uint64) entity.Categories {
	var category entity.Categories
	return category
}

func (db *categoryConnection) DeleteCategory(category entity.Categories) {
	db.connection.Delete(&category)
}

func (db *categoryConnection) AllCategory() []entity.Categories {
	var category []entity.Categories
	return category
}
