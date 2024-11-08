package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"github.com/DiegoJCordeiro/golang-study/chapter10/graph/model"
)

// Courses is the resolver for the courses field.
func (r *categoryResolver) Courses(_ context.Context, obj *model.Category) ([]*model.Course, error) {
	var coursesModel []*model.Course
	courses, err := r.CourseDB.FindAllCourseByCategoryID(obj.ID)

	if err != nil {
		return nil, err
	}

	for _, course := range courses {

		coursesModel = append(coursesModel, &model.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
		})
	}

	return coursesModel, nil
}

// Category is the resolver for the category field.
func (r *courseResolver) Category(_ context.Context, obj *model.Course) (*model.Category, error) {
	categoryModel, err := r.CategoryDB.FindCategoryByCourseID(obj.ID)

	if err != nil {
		return nil, err
	}

	categoryFound := &model.Category{
		ID:          categoryModel.ID,
		Name:        categoryModel.Name,
		Description: &categoryModel.Description,
	}

	return categoryFound, nil
}

// RegisterCategory is the resolver for the registerCategory field.
func (r *mutationResolver) RegisterCategory(_ context.Context, input model.RegisterCategory) (*model.Category, error) {
	category, err := r.CategoryDB.RegisterCategory(input.Name, *input.Description)

	if err != nil {
		return nil, err
	}

	return &model.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: &category.Description,
	}, nil
}

// RegisterCourse is the resolver for the registerCourse field.
func (r *mutationResolver) RegisterCourse(_ context.Context, input model.RegisterCourse) (*model.Course, error) {
	course, err := r.CourseDB.RegisterCourse(input.Name, *input.Description, input.CategoryID)

	if err != nil {
		return nil, err
	}

	return &model.Course{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
	}, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(_ context.Context) ([]*model.Category, error) {
	categories, err := r.CategoryDB.FindAllCategory()
	var categoriesModel []*model.Category

	if err != nil {
		return nil, err
	}

	for _, category := range categories {

		categoriesModel = append(categoriesModel, &model.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: &category.Description,
		})
	}

	return categoriesModel, nil
}

// Courses is the resolver for the courses field.
func (r *queryResolver) Courses(_ context.Context) ([]*model.Course, error) {
	courses, err := r.CourseDB.FindAllCourse()
	var coursesModel []*model.Course

	if err != nil {
		return nil, err
	}

	for _, course := range courses {
		coursesModel = append(coursesModel, &model.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: course.Description,
		})
	}

	return coursesModel, nil
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Course returns CourseResolver implementation.
func (r *Resolver) Course() CourseResolver { return &courseResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type courseResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	func (r *categoryResolver) Categories(_ context.Context, obj *model.Course) ([]*model.Category, error) {

	var categoriesModel []*model.Category
	categories, err := r.CategoryDB.FindAllCategoryByCourseID(obj.ID)

	if err != nil {
		return nil, err
	}

	for _, category := range categories {

		categoriesModel = append(categoriesModel, &model.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: &category.Description,
		})
	}

	return categoriesModel, nil
}
*/
