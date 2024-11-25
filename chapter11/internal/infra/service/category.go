package service

import (
	"context"
	"github.com/DiegoJCordeiro/golang-study/chapter11/internal/infra/database"
	"github.com/DiegoJCordeiro/golang-study/chapter11/internal/infra/proto"
	"google.golang.org/grpc/status"
)

type CategoryService struct {
	proto.UnimplementedCategoryServiceServer
	DB *database.CategoryDB
}

func NewCategoryService(db *database.CategoryDB) *CategoryService {
	return &CategoryService{
		DB: db,
	}
}

func (service *CategoryService) CreateCategory(_ context.Context, in *proto.CreateCategoryRequest) (*proto.CategoryResponse, error) {
	category, err := service.DB.RegisterCategory(in.Name, in.Description)

	if err != nil {
		return nil, err
	}

	return &proto.CategoryResponse{
		Category: &proto.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}, nil
}

func (service *CategoryService) UpdateCategory(_ context.Context, in *proto.UpdateCategoryRequest) (*proto.CategoryResponse, error) {

	category, errDB := service.DB.UpdateCategory(in.Id, in.Name, in.Description)

	if errDB != nil {
		return nil, errDB
	}

	return &proto.CategoryResponse{
		Category: &proto.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}, nil
}

func (service *CategoryService) FindAll(context.Context, *proto.Blank) (*proto.CategoryList, error) {

	categories := &proto.CategoryList{}

	categoriesFound, err := service.DB.FindAllCategory()

	if err != nil {
		return nil, err
	}

	for _, category := range categoriesFound {

		categories.Categories = append(categories.Categories, &proto.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return categories, nil
}

func (service *CategoryService) CreateCategoryInBatch(stream proto.CategoryService_CreateCategoryInBatchServer) error {

	categories := &proto.CategoryList{}

	for {

		received, err := stream.Recv()

		if err != nil {

			if status.Code(err) == 1 {
				return stream.Send(categories)
			}
			return err
		}

		category, errDB := service.DB.RegisterCategory(received.Name, received.Description)

		if errDB != nil {
			return errDB
		}

		categories.Categories = append(categories.Categories, &proto.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
}

func (service *CategoryService) DeleteCategoryInBatch(stream proto.CategoryService_DeleteCategoryInBatchServer) error {

	categories := &proto.CategoryList{}

	for {

		received, errStream := stream.Recv()

		if errStream != nil {
			if status.Code(errStream) == 1 {
				return stream.Send(categories)
			}

			return errStream
		}

		categoryFound, errDB := service.DB.FindCategoryByID(received.Id)

		if errDB != nil {
			return errDB
		}

		errDB = service.DB.DeleteCategoryById(received.Id)

		if errDB != nil {
			return errDB
		}

		categories.Categories = append(categories.Categories, &proto.Category{
			Id:          categoryFound.ID,
			Name:        categoryFound.Name,
			Description: categoryFound.Description,
		})
	}
}
