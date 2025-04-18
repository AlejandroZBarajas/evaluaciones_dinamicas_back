package categoryApplication

import (
	categoryDomain "evaluaciones/src/categories/domain"
	categoryEntity "evaluaciones/src/categories/domain/entity"
)

type UpdateCategory struct {
	repo categoryDomain.CategoryInterface
}

func NewUpdateCategory(repo categoryDomain.CategoryInterface) *UpdateCategory {
	return &UpdateCategory{repo: repo}
}

func (uc *UpdateCategory) Run(categoryID int32, category *categoryEntity.CategoryEntity) error {
	return uc.repo.UpdateCategory(categoryID, category)
}
