package repository

import (
	"app/infra/database"
	"app/usecases"
)

type CategoryRepositoryDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func NewCategoryRepository(
	sqlAdapter *database.SqlAdapter,
) usecases.CategoryRepository {
	return &CategoryRepositoryDependencies{sqlAdapter}
}
