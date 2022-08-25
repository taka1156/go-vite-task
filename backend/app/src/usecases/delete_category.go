package usecases

type DeleteCategoryAdapter interface {
	Do(categoryId int) (*bool, error)
}

type DeleteCategoryDependencies struct {
	deleteCategory DeleteCategoryAdapter
}

func (dep DeleteCategoryDependencies) Handle(categoryId int) (*bool, error) {
	isDeleted, err := dep.deleteCategory.Do(categoryId)
	if err != nil {
		return nil, err
	}

	return isDeleted, nil
}

func NewDeleteCategoryUsecase(deleteCategory DeleteCategoryAdapter) DeleteCategoryInterface {
	return &DeleteCategoryDependencies{deleteCategory}
}
