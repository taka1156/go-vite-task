package usecases

import (
	"app/entity/model"
)

type CreateCategoryAdapter interface {
	Do(input model.InputCategory) (*uint, error)
}

type CreateCategoryDependencies struct {
	updateCategory CreateCategoryAdapter
	getCategory    GetCategoryAdapter
}

func (dep CreateCategoryDependencies) Handle(updateInput model.InputCategory) (*model.Category, error) {
	categoryId, err := dep.updateCategory.Do(updateInput)
	if err != nil {
		return nil, err
	}

	team, err := dep.getCategory.Do(categoryId)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func NewCreateCategoryUsecase(createCategory CreateCategoryAdapter, getCategory GetCategoryAdapter) CreateCategoryInterface {
	return &CreateCategoryDependencies{createCategory, getCategory}
}
