package graph

import "github.com/DiegoJCordeiro/golang-study/chapter10/internal/infra/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB *database.CategoryDB
	CourseDB   *database.CourseDB
}
