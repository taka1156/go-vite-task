package usecases

import "app/entity/model"

type GetCategoriesAdapter interface {
	Do() ([]*model.Category, error)
}

type GetCategoriesDependencies struct {
	getCategories GetCategoriesAdapter
}

func (dep GetCategoriesDependencies) Handle() ([]*model.Category, error) {
	isDeleted, err := dep.getCategories.Do()
	if err != nil {
		return nil, err
	}

	return isDeleted, nil
}

func NewGetCategoriesUsecase(getCategories GetCategoriesAdapter) GetCategoriesInterface {
	return &GetCategoriesDependencies{getCategories}
}
