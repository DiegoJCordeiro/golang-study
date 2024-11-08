package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type CourseDB struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourseDB(db *sql.DB) *CourseDB {
	return &CourseDB{
		db: db,
	}
}

func (courseDB *CourseDB) RegisterCourse(name, description, categoryID string) (*CourseDB, error) {

	id := uuid.New().String()

	_, err := courseDB.db.Exec("INSERT INTO Courses (ID, Name, Description, Category_Id) VALUES ($1, $2, $3, $4)", id, name, description, categoryID)

	if err != nil {
		return nil, err
	}

	return &CourseDB{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil
}

func (courseDB *CourseDB) FindAllCourse() ([]CourseDB, error) {

	var courses []CourseDB

	rows, err := courseDB.db.Query("SELECT ID, Name, Description, Category_Id FROM Courses")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id, name, description, categoryId string

		if err := rows.Scan(&id, &name, &description, &categoryId); err != nil {
			return nil, err
		}

		courses = append(courses, CourseDB{ID: id, Name: name, Description: description, CategoryID: categoryId})
	}

	return courses, nil
}

func (courseDB *CourseDB) FindAllCourseByCategoryID(categoryId string) ([]CourseDB, error) {

	var courses []CourseDB

	rows, err := courseDB.db.Query("SELECT ID, Name, Description, Category_Id FROM Courses WHERE Category_Id = $1", categoryId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id, name, description, categoryID string

		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}

		courses = append(courses, CourseDB{ID: id, Name: name, Description: description, CategoryID: categoryID})
	}

	return courses, nil
}
