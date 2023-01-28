package usecases

import "app/usecases"

type DeleteCategoryInteractorDependencies struct {
	categoryRepository usecases.CategoryRepository
}

func NewDeleteCategoryUsecase(categoryRepository usecases.CategoryRepository) usecases.DeleteCategoryUsecase {
	return &DeleteCategoryInteractorDependencies{categoryRepository}
}

func (interactor DeleteCategoryInteractorDependencies) Handle(categoryId int) (bool, error) {
	// isDeleted, err := interactor.DeleteCategory(categoryId)
	// if err != nil {
	// 	return false, err
	// }

	return true, nil
}
