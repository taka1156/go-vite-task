package usecases

import (
	"app/entity/model"
	"app/usecases"
)

type CreateCategoryInteractorDependencies struct {
	categoryRepository usecases.CategoryRepository
}

func NewCreateCategoryUsecase(categoryRepository usecases.CategoryRepository) usecases.CreateCategoryUsecase {
	return &CreateCategoryInteractorDependencies{categoryRepository}
}

func (interactor CreateCategoryInteractorDependencies) Handle(updateInput model.InputCategory) (*model.Category, error) {
	// categoryId, err := interactor.categoryRepository.CreateCategory(updateInput)
	// if err != nil {
	// 	return nil, err
	// }

	// team, err := interactor.categoryRepository.GetCategory(categoryId)
	// if err != nil {
	// 	return nil, err
	// }

	category := model.Category{}

	return &category, nil
}
