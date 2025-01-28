package repository

import (
	"database/sql"
	"github.com/DiegoJCordeiro/golang-study/chapter13/internal/db"
)

type CategoryRepositoryInterface interface {
	QueryOneCategory(name string) (*db.Category, error)
	QueryAllCategory() (*[]db.Category, error)
	InsertOneCategory(uuid, name, description string) error
}

type CategoryRepository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) (CategoryRepositoryInterface, error) {
	return &CategoryRepository{
		DB: db,
	}, nil
}

func (categoryRepository *CategoryRepository) QueryOneCategory(name string) (*db.Category, error) {
	return nil, nil
}

func (categoryRepository *CategoryRepository) QueryAllCategory() (*[]db.Category, error) {
	return nil, nil
}

func (categoryRepository *CategoryRepository) InsertOneCategory(uuid, name, description string) error {
	return nil
}
