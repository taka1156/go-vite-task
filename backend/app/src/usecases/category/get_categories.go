package usecases

import (
	"app/entity/model"
	"app/usecases"
)

type GetCategoriesInteractorDependencies struct {
	categoryRepository usecases.CategoryRepository
}

func NewGetCategoriesUsecase(categoryRepository usecases.CategoryRepository) usecases.GetCategoriesUsecase {
	return &GetCategoriesInteractorDependencies{categoryRepository}
}

func (interactor GetCategoriesInteractorDependencies) Handle() ([]*model.Category, error) {
	// err := interactor.categoryRepository.GetCategories()
	// if err != nil {
	// 	return false, err
	// }

	categories := []*model.Category{}

	return categories, nil
}
