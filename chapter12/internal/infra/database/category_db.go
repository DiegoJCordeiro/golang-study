package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type CategoryDB struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

func (category *CategoryDB) UpdateCategory(id, name, description string) (CategoryDB, error) {

	_, err := category.db.Exec(
		"UPDATE Categories SET name = ?, description = ? WHERE id = ?",
		name,
		description,
		id,
	)

	if err != nil {
		return CategoryDB{}, err
	}

	return CategoryDB{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (category *CategoryDB) RegisterCategory(name, description string) (CategoryDB, error) {

	id := uuid.New().String()

	_, err := category.db.Exec(
		"INSERT INTO Categories (id, name, description) VALUES ($1, $2, $3)",
		id,
		name,
		description,
	)

	if err != nil {
		return CategoryDB{}, err
	}

	return CategoryDB{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (categoryDB *CategoryDB) FindCategoryByID(categoryId string) (CategoryDB, error) {

	category := CategoryDB{}

	rows, errDB := categoryDB.db.Query("SELECT Id, Name, Description FROM Categories WHERE id = ?", categoryId)

	if errDB != nil {
		return CategoryDB{}, errDB
	}

	defer rows.Close()

	for rows.Next() {

		if errScan := rows.Scan(&category.ID, &category.Name, &category.Description); errScan != nil {
			return CategoryDB{}, errScan
		}

		return category, nil
	}

	return CategoryDB{}, nil
}

func (category *CategoryDB) FindAllCategory() ([]CategoryDB, error) {

	categories := []CategoryDB{}

	rows, err := category.db.Query("SELECT id, name, description FROM Categories")

	if err != nil {
		return categories, err
	}

	defer rows.Close()

	for rows.Next() {
		var id, name, description string

		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}

		categories = append(categories, CategoryDB{ID: id, Name: name, Description: description})
	}

	return categories, nil
}

func (category *CategoryDB) FindCategoryByCourseID(courseId string) (*CategoryDB, error) {

	var categoryModel CategoryDB

	rows, err := category.db.Query("SELECT categories.ID, categories.Name, categories.Description FROM Categories categories JOIN Courses courses ON courses.id = $1 AND categories.id = courses.category_id", courseId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id, name, description string

		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}

		categoryModel = CategoryDB{ID: id, Name: name, Description: description}
	}

	return &categoryModel, nil
}

func (category *CategoryDB) DeleteCategoryById(categoryId string) error {
	_, err := category.db.Exec("DELETE FROM Categories WHERE id = $1", categoryId)
	if err != nil {
		return err
	}
	return nil
}
