package service

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/trisatya23/go-restful-api-master/exception"
	"github.com/trisatya23/go-restful-api-master/helper"
	"github.com/trisatya23/go-restful-api-master/model/domain"
	"github.com/trisatya23/go-restful-api-master/model/web"
	"github.com/trisatya23/go-restful-api-master/repository"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		Validate:           validate,
	}
}

// Create Category
func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) (web.CategoryResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return web.CategoryResponse{}, err
	}

	category := domain.Category{Name: request.Name}
	savedCategory, err := service.CategoryRepository.Save(ctx, category)
	if err != nil {
		return web.CategoryResponse{}, err
	}

	return helper.ToCategoryResponse(savedCategory), nil
}

// Update Category
func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) (web.CategoryResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return web.CategoryResponse{}, err
	}

	category, err := service.CategoryRepository.FindById(ctx, request.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.CategoryResponse{}, exception.NewNotFoundError("Category not found")
	} else if err != nil {
		return web.CategoryResponse{}, err
	}

	category.Name = request.Name
	updatedCategory, err := service.CategoryRepository.Update(ctx, category)
	if err != nil {
		return web.CategoryResponse{}, err
	}

	return helper.ToCategoryResponse(updatedCategory), nil
}

// Delete Category
func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) error {
	category, err := service.CategoryRepository.FindById(ctx, categoryId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return exception.NewNotFoundError("Category not found")
	} else if err != nil {
		return err
	}

	return service.CategoryRepository.Delete(ctx, category)
}

// Find Category By ID
func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) (web.CategoryResponse, error) {
	category, err := service.CategoryRepository.FindById(ctx, categoryId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.CategoryResponse{}, exception.NewNotFoundError("Category not found")
	} else if err != nil {
		return web.CategoryResponse{}, err
	}

	return helper.ToCategoryResponse(category), nil
}

// Find All Categories
func (service *CategoryServiceImpl) FindAll(ctx context.Context) ([]web.CategoryResponse, error) {
	categories, err := service.CategoryRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return helper.ToCategoryResponses(categories), nil
}
