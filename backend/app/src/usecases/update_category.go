package usecases

import (
	"app/entity/model"
)

type UpdateCategoryAdapter interface {
	Do(updateInput model.UpdateCategory) (*uint, error)
}

type UpdateCategoryDependencies struct {
	updateCategory UpdateCategoryAdapter
	getTeam        GetCategoryAdapter
}

func (dep UpdateCategoryDependencies) Handle(updateInput model.UpdateCategory) (*model.Category, error) {
	categoryId, err := dep.updateCategory.Do(updateInput)
	if err != nil {
		return nil, err
	}

	team, err := dep.getTeam.Do(categoryId)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func NewUpdateCategoryUsecase(updateCategory UpdateCategoryAdapter, getCategory GetCategoryAdapter) UpdateCategoryInterface {
	return &UpdateCategoryDependencies{updateCategory, getCategory}
}
