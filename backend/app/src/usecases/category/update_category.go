package usecases

import (
	"app/entity/model"
	"app/usecases"
)

type UpdateCategoryInteractorDependencies struct {
	categoryRepository usecases.CategoryRepository
}

func NewUpdateCategoryUsecase(categoryRepository usecases.CategoryRepository) usecases.UpdateCategoryUsecase {
	return &UpdateCategoryInteractorDependencies{categoryRepository}
}

func (interactor UpdateCategoryInteractorDependencies) Handle(updateInput model.UpdateCategory) (*model.Category, error) {
	// categoryId, err := interactor.categoryRepository.UpdateCategory(updateInput)
	// if err != nil {
	// 	return nil, err
	// }

	// team, err := interactor.categoryRepository.GetTeam(categoryId)
	// if err != nil {
	// 	return nil, err
	// }

	category := model.Category{}

	return &category, nil
}
