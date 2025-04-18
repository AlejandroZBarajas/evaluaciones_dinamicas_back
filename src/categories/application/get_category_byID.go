package categoryApplication

import (
	categoryDomain "evaluaciones/src/categories/domain"
	categoryEntity "evaluaciones/src/categories/domain/entity"
)

type GetCategoryByID struct {
	repo categoryDomain.CategoryInterface
}

func NewGetCategoryByID(repo categoryDomain.CategoryInterface) *GetCategoryByID {
	return &GetCategoryByID{repo: repo}
}

func (gc *GetCategoryByID) Run(id int32) (*categoryEntity.CategoryEntity, error) {
	return gc.repo.GetCategoryByID(id)
}
