package categoryApplication

import (
	categoryDomain "evaluaciones/src/categories/domain"
	categoryEntity "evaluaciones/src/categories/domain/entity"
)

type CreateCategory struct {
	repo categoryDomain.CategoryInterface
}

func NewCreateCategory(repo categoryDomain.CategoryInterface) *CreateCategory {
	return &CreateCategory{repo: repo}
}

func (cc *CreateCategory) Run(name string, teacherID int32) error {
	category := categoryEntity.CreateCategory(name, teacherID)
	return cc.repo.CreateCategory(category)
}
